package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
 	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func Test_Network(t *testing.T) {
	t.Parallel()

	uniqueId := random.UniqueId()
	nameNetworking := fmt.Sprintf("test-azure-mod-network-%s", uniqueId)
	nameWeb := fmt.Sprintf("test-azure-mod-web-%s", uniqueId)

 	terraformDirNetworking := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/networking")
 	terraformDirWebserver := test_structure.CopyTerraformFolderToTemp(t, "../", "examples/webserver")

	terraformOptionsNetworking := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		MaxRetries:         3,
		TimeBetweenRetries: 60 * time.Second,
		TerraformDir:       terraformDirNetworking,
		Vars: map[string]interface{}{
			"name": nameNetworking,
		},
	})
	defer terraform.Destroy(t, terraformOptionsNetworking)
	terraform.InitAndApply(t, terraformOptionsNetworking)

	outputNetworking := terraform.OutputAll(t, terraformOptionsNetworking)

    terraformOptionsWebserver := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    		MaxRetries:         3,
    		TimeBetweenRetries: 60 * time.Second,
    		TerraformDir:       terraformDirWebserver,
    		Vars: map[string]interface{}{
    			"name": nameWeb,
    			"vm_subnet_id": outputNetworking["vm_subnet_id"],
    		},
    	})
    	defer terraform.Destroy(t, terraformOptionsWebserver)
    	terraform.InitAndApply(t, terraformOptionsWebserver)

    	// Run `terraform output` to get the IP of the instance
        publicIp := terraform.Output(t, terraformOptionsWebserver, "public_ip")

        // Make an HTTP request to the instance and make sure we get back a 200 OK with the body "Hello, World!"
        url := fmt.Sprintf("http://%s:80", publicIp)
        http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, World!", 30, 5*time.Second)
}
