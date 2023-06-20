terraform {
  required_providers {
    usercreation = {
      version = "0.1"
      source = "local.com/rupeshshewalkar/usercreation"
    }
  }
}
resource "usercreation_user" "new" {
  name = "My User"
  email = "myuser@mail.com"
  username = "myuser"
  alias = "myuser.com"
}
