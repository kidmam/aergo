/**
 * @file    preprocess.h
 * @copyright defined in aergo/LICENSE.txt
 */

#ifndef _PREPROCESS_H
#define _PREPROCESS_H

#include "common.h"

#include "compile.h"

#ifndef _STRBUF_T
#define _STRBUF_T
typedef struct strbuf_s strbuf_t;
#endif  /* _STRBUF_T */

typedef struct subst_s {
    char *path;
    int line;

    int len;
    char *src;

    strbuf_t *res;
} subst_t;

int preprocess(char *path, opt_t opt, strbuf_t *res);

void append_directive(char *path, int line, strbuf_t *res);

#endif /*_PREPROCESS_H */
