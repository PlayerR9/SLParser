source : sourceRule source1 EOF ;
sourceRule : SOURCE_ID source1 EOF_ID ;
source1 : rule ;
source1 : rule source1 ;
rule : LOWERCASE_ID COLON rhs1 SEMICOLON ;
rule : LOWERCASE_ID COLON rhs1 rule1 ;
rule1 : PIPE rhs1 ;
rule1 : PIPE rhs1 rule1 ;
rhs1 : rhs ;
rhs1 : rhs rhs1 ;
rhs : identifier ;
rhs : OP_PAREN orExpr CL_PAREN ;
orExpr : identifier orExpr1 ;
orExpr1 : PIPE identifier ;
orExpr1 : PIPE identifier orExpr1 ;
identifier : UPPERCASE_ID ;
identifier : LOWERCASE_ID ;