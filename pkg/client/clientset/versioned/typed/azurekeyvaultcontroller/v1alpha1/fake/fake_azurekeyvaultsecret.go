/*
Copyright Sparebanken Vest

Authored by The Cumulus Team.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/SparebankenVest/azure-keyvault-controller/pkg/apis/azurekeyvaultcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureKeyVaultSecrets implements AzureKeyVaultSecretInterface
type FakeAzureKeyVaultSecrets struct {
	Fake *FakeAzurekeyvaultcontrollerV1alpha1
	ns   string
}

var azurekeyvaultsecretsResource = schema.GroupVersionResource{Group: "azurekeyvaultcontroller.spv.no", Version: "v1alpha1", Resource: "azurekeyvaultsecrets"}

var azurekeyvaultsecretsKind = schema.GroupVersionKind{Group: "azurekeyvaultcontroller.spv.no", Version: "v1alpha1", Kind: "AzureKeyVaultSecret"}

// Get takes name of the azureKeyVaultSecret, and returns the corresponding azureKeyVaultSecret object, and an error if there is any.
func (c *FakeAzureKeyVaultSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azurekeyvaultsecretsResource, c.ns, name), &v1alpha1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultSecret), err
}

// List takes label and field selectors, and returns the list of AzureKeyVaultSecrets that match those selectors.
func (c *FakeAzureKeyVaultSecrets) List(opts v1.ListOptions) (result *v1alpha1.AzureKeyVaultSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azurekeyvaultsecretsResource, azurekeyvaultsecretsKind, c.ns, opts), &v1alpha1.AzureKeyVaultSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureKeyVaultSecretList{ListMeta: obj.(*v1alpha1.AzureKeyVaultSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureKeyVaultSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureKeyVaultSecrets.
func (c *FakeAzureKeyVaultSecrets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azurekeyvaultsecretsResource, c.ns, opts))

}

// Create takes the representation of a azureKeyVaultSecret and creates it.  Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultSecrets) Create(azureKeyVaultSecret *v1alpha1.AzureKeyVaultSecret) (result *v1alpha1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azurekeyvaultsecretsResource, c.ns, azureKeyVaultSecret), &v1alpha1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultSecret), err
}

// Update takes the representation of a azureKeyVaultSecret and updates it. Returns the server's representation of the azureKeyVaultSecret, and an error, if there is any.
func (c *FakeAzureKeyVaultSecrets) Update(azureKeyVaultSecret *v1alpha1.AzureKeyVaultSecret) (result *v1alpha1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azurekeyvaultsecretsResource, c.ns, azureKeyVaultSecret), &v1alpha1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureKeyVaultSecrets) UpdateStatus(azureKeyVaultSecret *v1alpha1.AzureKeyVaultSecret) (*v1alpha1.AzureKeyVaultSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azurekeyvaultsecretsResource, "status", c.ns, azureKeyVaultSecret), &v1alpha1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultSecret), err
}

// Delete takes name of the azureKeyVaultSecret and deletes it. Returns an error if one occurs.
func (c *FakeAzureKeyVaultSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azurekeyvaultsecretsResource, c.ns, name), &v1alpha1.AzureKeyVaultSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureKeyVaultSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azurekeyvaultsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureKeyVaultSecretList{})
	return err
}

// Patch applies the patch and returns the patched azureKeyVaultSecret.
func (c *FakeAzureKeyVaultSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AzureKeyVaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azurekeyvaultsecretsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureKeyVaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureKeyVaultSecret), err
}
