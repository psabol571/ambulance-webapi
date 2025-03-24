/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xsabol@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ambulance_wl

import (
	"github.com/gin-gonic/gin"
)

type AmbulanceConditionsAPI interface {


    // GetConditions Get /api/waiting-list/:ambulanceId/condition
    // Provides the list of conditions associated with ambulance 
     GetConditions(c *gin.Context)

}