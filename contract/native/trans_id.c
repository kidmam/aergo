/**
 * @file    trans_id.c
 * @copyright defined in aergo/LICENSE.txt
 */

#include "common.h"

#include "array.h"
#include "ir_fn.h"
#include "ir_bb.h"
#include "trans_blk.h"
#include "trans_stmt.h"
#include "trans_exp.h"

#include "trans_id.h"

static void
id_trans_var(trans_t *trans, ast_id_t *id)
{
    if (is_global_id(id)) {
        /* Initialization of the global variable will be done in the constructor */
        ir_add_global(trans->ir, id);
        return;
    }

    ASSERT(trans->fn != NULL);

    if (is_stack_id(id))
        fn_add_stack(trans->fn, id);
    else
        fn_add_local(trans->fn, id);
}

static ast_exp_t *
gen_dflt_exp(meta_t *meta)
{
    ast_exp_t *dflt_exp = exp_new_lit(&meta->pos);

    if (is_array_type(meta)) {
        value_set_ptr(&dflt_exp->u_lit.val, xcalloc(meta->arr_size), meta->arr_size);
    }
    else if (is_bool_type(meta)) {
        value_set_bool(&dflt_exp->u_lit.val, false);
    }
    else if (is_fpoint_type(meta)) {
        value_set_f64(&dflt_exp->u_lit.val, 0.0);
    }
    else if (is_integer_type(meta) || is_pointer_type(meta)) {
        value_set_i64(&dflt_exp->u_lit.val, 0);
    }
    else {
        ASSERT1(is_struct_type(meta), meta->type);
        value_set_ptr(&dflt_exp->u_lit.val, xcalloc(meta_size(meta)), meta_size(meta));
    }

    meta_copy(&dflt_exp->meta, meta);
    meta_set_undef(&dflt_exp->meta);

    return dflt_exp;
}

static void
gen_init_stmt(trans_t *trans, ast_id_t *id)
{
    ast_exp_t *dflt_exp, *id_exp;

    ASSERT1(is_global_id(id), id->up->kind);

    if (id->u_var.dflt_exp == NULL)
        id->u_var.dflt_exp = gen_dflt_exp(&id->meta);

    dflt_exp = id->u_var.dflt_exp;

    id_exp = exp_new_id_ref(id->name, &dflt_exp->pos);

    id_exp->id = id;
    meta_copy(&id_exp->meta, &id->meta);

    ASSERT2(meta_cmp(&id_exp->meta, &dflt_exp->meta) == 0, id_exp->meta.type,
            dflt_exp->meta.type);

    stmt_trans(trans, stmt_new_assign(id_exp, dflt_exp, &dflt_exp->pos));
}

static void
id_trans_fn(trans_t *trans, ast_id_t *id)
{
    ir_fn_t *fn = fn_new(id);

    trans->fn = fn;
    trans->bb = fn->entry_bb;

    if (is_ctor_id(id)) {
        int i;
        ast_id_t *cont_id = id->up;

        ASSERT(cont_id != NULL);
        ASSERT1(is_cont_id(cont_id), cont_id->kind);

        array_foreach(&cont_id->u_cont.blk->ids, i) {
            ast_id_t *fld_id = array_get_id(&cont_id->u_cont.blk->ids, i);

            if (is_var_id(fld_id)) {
                gen_init_stmt(trans, fld_id);
            }
            else if (is_tuple_id(fld_id)) {
                int i;

                array_foreach(fld_id->u_tup.elem_ids, i) {
                    gen_init_stmt(trans, array_get_id(fld_id->u_tup.elem_ids, i));
                }
            }
        }

        ASSERT(trans->bb == fn->entry_bb);
    }

    if (id->u_fn.blk != NULL) {
        blk_trans(trans, id->u_fn.blk);

        if (trans->bb != NULL) {
            bb_add_branch(trans->bb, NULL, fn->exit_bb);
            fn_add_basic_blk(fn, trans->bb);
        }

        fn_add_basic_blk(fn, fn->exit_bb);
    }
    else {
        fn_add_basic_blk(fn, fn->entry_bb);
    }

    trans->fn = NULL;
    trans->bb = NULL;

    ir_add_fn(trans->ir, fn);
}

static void
id_trans_contract(trans_t *trans, ast_id_t *id)
{
    ASSERT(id->u_cont.blk != NULL);

    blk_trans(trans, id->u_cont.blk);
}

static void
id_trans_label(trans_t *trans, ast_id_t *id)
{
    id->u_lab.stmt->label_bb = bb_new();
}

static void
id_trans_tuple(trans_t *trans, ast_id_t *id)
{
    int i;

    array_foreach(id->u_tup.elem_ids, i) {
        id_trans_var(trans, array_get_id(id->u_tup.elem_ids, i));
    }
}

void
id_trans(trans_t *trans, ast_id_t *id)
{
    switch (id->kind) {
    case ID_VAR:
        id_trans_var(trans, id);
        break;

    case ID_FN:
        id_trans_fn(trans, id);
        break;

    case ID_CONT:
        id_trans_contract(trans, id);
        break;

    case ID_LABEL:
        id_trans_label(trans, id);
        break;

    case ID_TUPLE:
        id_trans_tuple(trans, id);
        break;

    case ID_STRUCT:
    case ID_ENUM:
    case ID_ITF:
        break;

    default:
        ASSERT1(!"invalid identifier", id->kind);
    }
}

/* end of trans_id.c */