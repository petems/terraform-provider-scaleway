package scaleway

import (
  "github.com/hashicorp/terraform/helper/schema"
  api "github.com/nicolai86/scaleway-sdk"
)

func resourceScalewaySSHKeys() *schema.Resource {
 return &schema.Resource{
    Create: resourceScalewaySSHKeysCreate,
    Read:   resourceScalewaySSHKeysRead,
    Update: resourceScalewaySSHKeysUpdate,
    Delete: resourceScalewaySSHKeysDelete,

    Schema: map[string]*schema.Schema{
      "user_id": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
      "ssh_public_keys": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
 }
}

func resourceScalewaySSHKeysCreate(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceScalewaySSHKeysRead(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceScalewaySSHKeysUpdate(d *schema.ResourceData, m interface{}) error {
  return nil
}

func resourceScalewaySSHKeysDelete(d *schema.ResourceData, m interface{}) error {
  return nil
}
