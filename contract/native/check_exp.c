/**
 * @file    check_exp.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "ast_id.h"
#include "ast_blk.h"
#include "check_id.h"

#include "check_exp.h"

static bool
exp_check_id_ref(check_t *check, ast_exp_t *exp)
{
    ast_id_t *id = NULL;
    char *name = exp->u_id.name;

    ASSERT1(is_id_ref_exp(exp), exp->kind);
    ASSERT(name != NULL);

    if (strcmp(name, "this") == 0) {
        id = check->cont_id;
    }
    else if (check->qual_id != NULL) {
        id = id_lookup_fld(check->qual_id, name, check->cont_id == check->qual_id);
    }
    else {
        if (check->fn_id != NULL)
            id = id_lookup_param(check->fn_id, name);

        if (id == NULL)
            id = blk_lookup_id(check->blk, name, exp->num);
    }

    if (id == NULL)
        RETURN(ERROR_UNDEFINED_ID, &exp->pos, name);

    id_trycheck(check, id);

    id->is_used = true;

    if (is_const_id(id) && id->val != NULL) {
        exp->kind = EXP_LIT;
        exp->u_lit.val = *id->val;
    }
    else {
        exp->id = id;
    }

    meta_copy(&exp->meta, &id->meta);

    return true;
}

static bool
exp_check_lit(check_t *check, ast_exp_t *exp)
{
    ASSERT1(is_lit_exp(exp), exp->kind);

    switch (exp->u_lit.val.type) {
    case TYPE_BOOL:
        meta_set_bool(&exp->meta);
        break;
    case TYPE_UINT64:
        meta_set_uint64(&exp->meta);
        meta_set_undef(&exp->meta);
        break;
    case TYPE_DOUBLE:
        meta_set_double(&exp->meta);
        meta_set_undef(&exp->meta);
        break;
    case TYPE_STRING:
        meta_set_string(&exp->meta);
        break;
    case TYPE_OBJECT:
        meta_set_object(&exp->meta, NULL);
        meta_set_undef(&exp->meta);
        break;
    default:
        ASSERT1(!"invalid value", exp->u_lit.val.type);
    }

    return true;
}

static bool
exp_check_array(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *id_exp;
    meta_t *id_meta;
    ast_exp_t *idx_exp;
    meta_t *idx_meta;

    ASSERT1(is_array_exp(exp), exp->kind);
    ASSERT(exp->u_arr.id_exp != NULL);

    id_exp = exp->u_arr.id_exp;
    id_meta = &id_exp->meta;

    CHECK(exp_check(check, id_exp));

    exp->id = id_exp->id;
    ASSERT(exp->id != NULL);

    idx_exp = exp->u_arr.idx_exp;
    idx_meta = &idx_exp->meta;

    CHECK(exp_check(check, idx_exp));

    if (is_array_type(id_meta)) {
        ASSERT(id_meta->arr_dim > 0);
        ASSERT(id_meta->dim_sizes != NULL);

        if (!is_integer_type(idx_meta))
            RETURN(ERROR_INVALID_SIZE_VAL, &idx_exp->pos, meta_to_str(idx_meta));

        meta_copy(&exp->meta, id_meta);

        if (is_lit_exp(idx_exp)) {
            ASSERT(id_meta->dim_sizes != NULL);

            /* dim_sizes[0] can be negative if array is used as a parameter */
            if (id_meta->dim_sizes[0] > 0 &&
                val_i64(&idx_exp->u_lit.val) >= (uint)id_meta->dim_sizes[0])
                RETURN(ERROR_INVALID_ARR_IDX, &idx_exp->pos);
        }

        /* Whenever an array element is accessed, strip it by one dimension */
        meta_strip_arr_dim(&exp->meta);
    }
    else {
        if (!is_map_type(id_meta))
            RETURN(ERROR_INVALID_SUBSCRIPT, &id_exp->pos);

        CHECK(meta_cmp(id_meta->elems[0], idx_meta));

        meta_eval(id_meta->elems[0], idx_meta);
        meta_copy(&exp->meta, id_meta->elems[1]);
    }

    return true;
}

static bool
exp_check_cast(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *val_exp;
    meta_t *val_meta;

    ASSERT1(is_cast_exp(exp), exp->kind);
    ASSERT1(is_valid_type(exp->u_cast.to_meta.type), exp->u_cast.to_meta.type);
    ASSERT(exp->u_cast.val_exp != NULL);

    val_exp = exp->u_cast.val_exp;
    val_meta = &val_exp->meta;

    CHECK(exp_check(check, val_exp));

    meta_copy(&exp->meta, &exp->u_cast.to_meta);

    if (is_array_type(val_meta) || !is_compatible_type(&exp->meta, val_meta))
        RETURN(ERROR_INCOMPATIBLE_TYPE, &val_exp->pos, meta_to_str(val_meta),
               meta_to_str(&exp->meta));

    if (is_lit_exp(val_exp)) {
        exp->u_lit.val = val_exp->u_lit.val;
        value_cast(&exp->u_lit.val, &exp->meta);

        exp->kind = EXP_LIT;
        meta_set_undef(&exp->meta);
    }

    return true;
}

static bool
exp_check_unary(check_t *check, ast_exp_t *exp)
{
    op_kind_t op = exp->u_un.kind;
    ast_exp_t *val_exp;
    meta_t *val_meta;

    ASSERT1(is_unary_exp(exp), exp->kind);
    ASSERT(exp->u_un.val_exp != NULL);

    val_exp = exp->u_un.val_exp;
    val_meta = &val_exp->meta;

    CHECK(exp_check(check, val_exp));

    switch (op) {
    case OP_INC:
    case OP_DEC:
        if (!is_usable_lval(val_exp))
            RETURN(ERROR_INVALID_LVALUE, &val_exp->pos);

        if (!is_integer_type(val_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &val_exp->pos, meta_to_str(val_meta));

        meta_copy(&exp->meta, val_meta);
        break;

    case OP_NEG:
        if (!is_numeric_type(val_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &val_exp->pos, meta_to_str(val_meta));

        meta_copy(&exp->meta, val_meta);
        break;

    case OP_NOT:
        if (!is_bool_type(val_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &val_exp->pos, meta_to_str(val_meta));

        meta_copy(&exp->meta, val_meta);
        break;

    default:
        ASSERT1(!"invalid operator", exp->u_un.kind);
    }

    if (is_lit_exp(val_exp)) {
        value_eval(op, &val_exp->u_lit.val, NULL, &exp->u_lit.val);

        exp->kind = EXP_LIT;
        meta_set_undef(&exp->meta);
    }

    return true;
}

static bool
exp_check_op_arith(check_t *check, ast_exp_t *exp)
{
    op_kind_t op = exp->u_bin.kind;
    ast_exp_t *l_exp, *r_exp;
    meta_t *l_meta, *r_meta;

    ASSERT(exp->u_bin.l_exp != NULL);
    ASSERT(exp->u_bin.r_exp != NULL);

    l_exp = exp->u_bin.l_exp;
    l_meta = &l_exp->meta;

    CHECK(exp_check(check, l_exp));

    if (op == OP_ADD) {
        if (!is_numeric_type(l_meta) && !is_string_type(l_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &l_exp->pos, meta_to_str(l_meta));
    }
    else if (op == OP_MOD) {
        if (!is_integer_type(l_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &l_exp->pos, meta_to_str(l_meta));
    }
    else if (!is_numeric_type(l_meta)) {
        RETURN(ERROR_INVALID_OP_TYPE, &l_exp->pos, meta_to_str(l_meta));
    }

    r_exp = exp->u_bin.r_exp;
    r_meta = &r_exp->meta;

    CHECK(exp_check(check, r_exp));
    CHECK(meta_cmp(l_meta, r_meta));

    meta_eval(l_meta, r_meta);
    meta_copy(&exp->meta, l_meta);

    return true;
}

static bool
exp_check_op_bit(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *l_exp, *r_exp;
    meta_t *l_meta, *r_meta;

    ASSERT(exp->u_bin.l_exp != NULL);
    ASSERT(exp->u_bin.r_exp != NULL);

    l_exp = exp->u_bin.l_exp;
    l_meta = &l_exp->meta;

    CHECK(exp_check(check, l_exp));

    if (!is_integer_type(l_meta))
        RETURN(ERROR_INVALID_OP_TYPE, &l_exp->pos, meta_to_str(l_meta));

    r_exp = exp->u_bin.r_exp;
    r_meta = &r_exp->meta;

    CHECK(exp_check(check, r_exp));

    switch (exp->u_bin.kind) {
    case OP_BIT_AND:
    case OP_BIT_OR:
    case OP_BIT_XOR:
        if (!is_integer_type(r_meta))
            RETURN(ERROR_INVALID_OP_TYPE, &r_exp->pos, meta_to_str(r_meta));
        break;

    case OP_RSHIFT:
    case OP_LSHIFT:
        if (!is_unsigned_type(r_meta) ||
            (is_lit_exp(r_exp) && is_signed_val(&r_exp->u_lit.val)))
            RETURN(ERROR_INVALID_OP_TYPE, &r_exp->pos, meta_to_str(r_meta));
        break;

    default:
        ASSERT1(!"invalid operator", exp->u_bin.kind);
    }

    meta_eval(l_meta, r_meta);
    meta_copy(&exp->meta, l_meta);

    return true;
}

static bool
exp_check_op_cmp(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *l_exp, *r_exp;
    meta_t *l_meta, *r_meta;

    ASSERT(exp->u_bin.l_exp != NULL);
    ASSERT(exp->u_bin.r_exp != NULL);

    l_exp = exp->u_bin.l_exp;
    l_meta = &l_exp->meta;

    CHECK(exp_check(check, l_exp));

    r_exp = exp->u_bin.r_exp;
    r_meta = &r_exp->meta;

    CHECK(exp_check(check, r_exp));

    if (is_tuple_type(l_meta))
        RETURN(ERROR_INVALID_OP_TYPE, &l_exp->pos, meta_to_str(l_meta));
    else if (is_tuple_type(r_meta))
        RETURN(ERROR_INVALID_OP_TYPE, &r_exp->pos, meta_to_str(r_meta));

    CHECK(meta_cmp(l_meta, r_meta));

    meta_eval(l_meta, r_meta);
    meta_set_bool(&exp->meta);

    return true;
}

static bool
exp_check_op_bool_cmp(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *l_exp, *r_exp;
    meta_t *l_meta, *r_meta;

    ASSERT(exp->u_bin.l_exp != NULL);
    ASSERT(exp->u_bin.r_exp != NULL);

    l_exp = exp->u_bin.l_exp;
    l_meta = &l_exp->meta;

    CHECK(exp_check(check, l_exp));

    if (!is_bool_type(l_meta))
        RETURN(ERROR_INVALID_COND_TYPE, &l_exp->pos, meta_to_str(l_meta));

    r_exp = exp->u_bin.r_exp;
    r_meta = &r_exp->meta;

    CHECK(exp_check(check, r_exp));

    if (!is_bool_type(r_meta))
        RETURN(ERROR_INVALID_COND_TYPE, &r_exp->pos, meta_to_str(r_meta));

    CHECK(meta_cmp(l_meta, r_meta));

    meta_eval(l_meta, r_meta);
    meta_set_bool(&exp->meta);

    return true;
}

static bool
exp_check_binary(check_t *check, ast_exp_t *exp)
{
    op_kind_t op = exp->u_bin.kind;

    ASSERT1(is_binary_exp(exp), exp->kind);

    switch (op) {
    case OP_ADD:
    case OP_SUB:
    case OP_MUL:
    case OP_DIV:
    case OP_MOD:
        CHECK(exp_check_op_arith(check, exp));
        break;

    case OP_BIT_AND:
    case OP_BIT_OR:
    case OP_BIT_XOR:
    case OP_RSHIFT:
    case OP_LSHIFT:
        CHECK(exp_check_op_bit(check, exp));
        break;

    case OP_EQ:
    case OP_NE:
    case OP_LT:
    case OP_GT:
    case OP_LE:
    case OP_GE:
        CHECK(exp_check_op_cmp(check, exp));
        break;

    case OP_AND:
    case OP_OR:
        CHECK(exp_check_op_bool_cmp(check, exp));
        break;

    default:
        ASSERT1(!"invalid operator", exp->u_bin.kind);
    }

    if (is_lit_exp(exp->u_bin.l_exp)) {
        ast_exp_t *l_exp = exp->u_bin.l_exp;
        ast_exp_t *r_exp = exp->u_bin.r_exp;

        if (!is_lit_exp(r_exp))
            return true;

        if ((op == OP_DIV || op == OP_MOD) && is_zero_val(&r_exp->u_lit.val))
            RETURN(ERROR_DIVIDE_BY_ZERO, &r_exp->pos);

        value_eval(op, &l_exp->u_lit.val, &r_exp->u_lit.val, &exp->u_lit.val);

        exp->kind = EXP_LIT;
        meta_set_undef(&exp->meta);
    }

    return true;
}

static bool
exp_check_ternary(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *pre_exp, *in_exp, *post_exp;
    meta_t *pre_meta, *in_meta, *post_meta;

    ASSERT1(is_ternary_exp(exp), exp->kind);
    ASSERT(exp->u_tern.pre_exp != NULL);
    ASSERT(exp->u_tern.in_exp != NULL);
    ASSERT(exp->u_tern.post_exp != NULL);

    pre_exp = exp->u_tern.pre_exp;
    pre_meta = &pre_exp->meta;

    CHECK(exp_check(check, pre_exp));

    if (!is_bool_type(pre_meta))
        RETURN(ERROR_INVALID_COND_TYPE, &pre_exp->pos, meta_to_str(pre_meta));

    in_exp = exp->u_tern.in_exp;
    in_meta = &in_exp->meta;

    CHECK(exp_check(check, in_exp));

    post_exp = exp->u_tern.post_exp;
    post_meta = &post_exp->meta;

    CHECK(exp_check(check, post_exp));
    CHECK(meta_cmp(in_meta, post_meta));

    meta_eval(in_meta, post_meta);
    meta_copy(&exp->meta, in_meta);

    return true;
}

static bool
exp_check_access(check_t *check, ast_exp_t *exp)
{
    ast_exp_t *id_exp, *fld_exp;
    meta_t *id_meta;
    ast_id_t *id;

    ASSERT1(is_access_exp(exp), exp->kind);

    id_exp = exp->u_acc.id_exp;
    id_meta = &id_exp->meta;

    CHECK(exp_check(check, id_exp));

    id = id_exp->id;

    if (id == NULL || is_tuple_type(id_meta) ||
        (is_fn_id(id) && !is_struct_type(id_meta) && !is_object_type(id_meta)))
        RETURN(ERROR_INACCESSIBLE_TYPE, &id_exp->pos, meta_to_str(id_meta));

    /* get the actual struct, contract or interface identifier */
    if (is_struct_type(id_meta) || is_object_type(id_meta)) {
        /* If "this.x" is used, the id is the contract identifier */
        ASSERT1(is_var_id(id) || is_fn_id(id) || is_cont_id(id), id->kind);
        ASSERT(id_meta->type_id != NULL);

        id = id_meta->type_id;
    }

    ASSERT(id != NULL);
    ASSERT1(is_type_id(id) || is_enum_id(id), id->kind);

    fld_exp = exp->u_acc.fld_exp;

    check->qual_id = id;

    if (exp_check(check, fld_exp)) {
        if (is_lit_exp(fld_exp)) {
            /* enum or contract constant */
            exp->kind = EXP_LIT;
            exp->u_lit.val = fld_exp->u_lit.val;
        }
        else {
            exp->id = fld_exp->id;
        }

        meta_copy(&exp->meta, &fld_exp->meta);
    }

    check->qual_id = NULL;

    return true;
}

static bool
exp_check_call(check_t *check, ast_exp_t *exp)
{
    int i;
    ast_exp_t *id_exp;
    array_t *param_exps;
    ast_id_t *id;
    array_t *param_ids;

    ASSERT1(is_call_exp(exp), exp->kind);

    id_exp = exp->u_call.id_exp;
    param_exps = exp->u_call.param_exps;

    if (is_id_ref_exp(id_exp) && strcmp(id_exp->u_id.name, "map") == 0) {
        /* In case of new map() */
        if (param_exps != NULL) {
            ast_exp_t *param_exp;

            ASSERT1(array_size(param_exps) == 1, array_size(param_exps));
            param_exp = array_get_exp(param_exps, 0);

            CHECK(exp_check(check, param_exp));
            ASSERT1(is_integer_type(&param_exp->meta), param_exp->meta.type);
        }

        meta_set_map(&exp->meta, NULL, NULL);
        meta_set_undef(&exp->meta);

        return true;
    }

    CHECK(exp_check(check, id_exp));

    id = id_exp->id;
    if (id == NULL)
        RETURN(ERROR_NOT_CALLABLE_EXP, &id_exp->pos);

    if (is_cont_id(id))
        /* lookup constructor */
        id = blk_lookup_fn(id->u_cont.blk, id->name);
    else if (is_var_id(id))
        /* lookup function */
        id = blk_lookup_fn(check->blk, id->name);

    if (!is_fn_id(id))
        RETURN(ERROR_NOT_CALLABLE_EXP, &id_exp->pos);

    param_ids = id->u_fn.param_ids;

    if (array_size(param_ids) != array_size(param_exps))
        RETURN(ERROR_MISMATCHED_COUNT, &id_exp->pos, "parameter", array_size(param_ids),
               array_size(param_exps));

    array_foreach(param_exps, i) {
        ast_id_t *param_id = array_get_id(param_ids, i);
        ast_exp_t *param_exp = array_get_exp(param_exps, i);

        CHECK(exp_check(check, param_exp));
        CHECK(meta_cmp(&param_id->meta, &param_exp->meta));

        meta_eval(&param_id->meta, &param_exp->meta);

        exp_check_overflow(param_exp, &param_id->meta);
    }

    exp->id = id;
    meta_copy(&exp->meta, &id->meta);

    return true;
}

static bool
exp_check_sql(check_t *check, ast_exp_t *exp)
{
    ASSERT1(is_sql_exp(exp), exp->kind);
    ASSERT(exp->u_sql.sql != NULL);

    switch (exp->u_sql.kind) {
    case SQL_QUERY:
        // TODO: need column meta
        break;

    case SQL_INSERT:
    case SQL_UPDATE:
    case SQL_DELETE:
        meta_set_int32(&exp->meta);
        break;

    default:
        ASSERT1(!"invalid sql", exp->u_sql.kind);
    }

    return true;
}

static bool
exp_check_tuple(check_t *check, ast_exp_t *exp)
{
    int i;
    array_t *elem_exps = exp->u_tup.elem_exps;

    ASSERT1(is_tuple_exp(exp), exp->kind);
    ASSERT(elem_exps != NULL);

    array_foreach(elem_exps, i) {
        CHECK(exp_check(check, array_get_exp(elem_exps, i)));
    }

    meta_set_tuple(&exp->meta, elem_exps);

    return true;
}

static bool
exp_check_init(check_t *check, ast_exp_t *exp)
{
    int i;
    bool is_aggr_lit = true;
    array_t *elem_exps = exp->u_init.elem_exps;

    ASSERT1(is_init_exp(exp), exp->kind);
    ASSERT(elem_exps != NULL);

    array_foreach(elem_exps, i) {
        ast_exp_t *elem_exp = array_get_exp(elem_exps, i);

        ASSERT1(!is_tuple_exp(elem_exp), elem_exp->kind);

        CHECK(exp_check(check, elem_exp));

        if (!is_lit_exp(elem_exp))
            is_aggr_lit = false;
    }

    meta_set_tuple(&exp->meta, elem_exps);

    if (is_aggr_lit) {
        int size = 0;
        char *raw = xcalloc(meta_size(&exp->meta));

        array_foreach(elem_exps, i) {
            ast_exp_t *elem_exp = array_get_exp(elem_exps, i);
            value_t *elem_val = &elem_exp->u_lit.val;

            size = ALIGN(size, meta_align(&elem_exp->meta));

            memcpy(raw + size, val_ptr(elem_val), val_size(elem_val));
            size += meta_size(&elem_exp->meta);
        }

        ASSERT2(ALIGN64(size) == meta_size(&exp->meta), size, meta_size(&exp->meta));

        exp->kind = EXP_LIT;
        value_set_ptr(&exp->u_lit.val, raw, size);
    }

    return true;
}

bool
exp_check(check_t *check, ast_exp_t *exp)
{
    ASSERT(exp != NULL);

    switch (exp->kind) {
    case EXP_NULL:
        return true;

    case EXP_ID_REF:
        return exp_check_id_ref(check, exp);

    case EXP_LIT:
        return exp_check_lit(check, exp);

    case EXP_ARRAY:
        return exp_check_array(check, exp);

    case EXP_CAST:
        return exp_check_cast(check, exp);

    case EXP_UNARY:
        return exp_check_unary(check, exp);

    case EXP_BINARY:
        return exp_check_binary(check, exp);

    case EXP_TERNARY:
        return exp_check_ternary(check, exp);

    case EXP_ACCESS:
        return exp_check_access(check, exp);

    case EXP_CALL:
        return exp_check_call(check, exp);

    case EXP_SQL:
        return exp_check_sql(check, exp);

    case EXP_TUPLE:
        return exp_check_tuple(check, exp);

    case EXP_INIT:
        return exp_check_init(check, exp);

    default:
        ASSERT1(!"invalid expression", exp->kind);
    }

    return true;
}

/* end of check_exp.c */
