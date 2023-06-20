package usercreation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strconv"
	"time"
)

const baseURL = "http://localhost:8000"

// User Create a struct to handler user object.
type User struct {
	ID       int    `json:"_id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	name := d.Get("name").(string)
	alias := d.Get("alias").(string)
	email := d.Get("email").(string)
	username := d.Get("username").(string)

	user := &User{
		Name:     name,
		Alias:    alias,
		Email:    email,
		Username: username,
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users/createNewUser", baseURL), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	var response User
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return err
	}

	d.Set("name", response.Name)
	d.Set("email", response.Email)
	d.Set("alias", response.Alias)
	d.Set("username", response.Username)
	d.SetId(strconv.Itoa(response.ID))

	return nil
}

// Function definition to read the resource.

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Function definition to update the resource.

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Function definition to delete the resource.

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Resource schema definition

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString, // Field type
				Computed: true,              // This flag means that the fields will be created after some processing
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, // Field is required
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
