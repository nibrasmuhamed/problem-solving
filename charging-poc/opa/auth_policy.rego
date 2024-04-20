package authz

allow {
    input.sub == "admin"
    input.obj == "showUserDetails"
    input.act == "POST"
}
