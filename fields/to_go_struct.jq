# cat fields/*.json | jq -f fields/to_go_struct.jq  -r | gofmt | goimports

def to_camel_case:
    ascii_downcase | sub("^(?<a>[a-z])"; .a|ascii_upcase) | gsub( "_(?<a>[a-z])"; .a|ascii_upcase)
    | gsub("Id$"; "ID")
    | gsub("Vat"; "VAT")
;

def field_type:
    if . == "C" then
        "string"
    elif . == "L" then
        "bool"
    elif . == "N" then
        "float64"
    elif . == "T" then
        "DateTime"
    elif . == "I" then
        "int"
    elif . == "Y" then
        "int"
    elif . == "M" then
        "string"
    else
        error("Unknown field type \(.)")
    end
;

# type Company struct {
# 	ID          string `json:"Id"`
# 	Description string `json:"Description"`
# 	Code        string `json:"Code"`
# }

(first | .DBF_NAME) as $struct_name
| (
    ["type \($struct_name) struct {"]
    +
    map(
        (.FIELD_NAME | to_camel_case) as $name
        | (.FIELD_TYPE |field_type) as $type
        | (.FIELD_NAME) as $json_name
        | (.FLD_DESC) as $comment
        | "\($name) \($type) `json:\"\($json_name)\"` // \($comment)"
    )
    +
    ["}"]
) | .[]
