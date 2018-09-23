package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

// ListClient is the an API for face detection, verification, and identification.
type ListClient struct {
	BaseClient
}

// NewListClient creates an instance of the ListClient client.
func NewListClient(azureRegion AzureRegions) ListClient {
	return ListClient{New(azureRegion)}
}

// AddFaceFromStream add a face to a face list. The input face is specified as an image with a targetFace rectangle. It
// returns a persistedFaceId representing the added face, and persistedFaceId will not expire.
// Parameters:
// faceListID - id referencing a particular face list.
// imageParameter - an image stream.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client ListClient) AddFaceFromStream(ctx context.Context, faceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "AddFaceFromStream", err.Error())
	}

	req, err := client.AddFaceFromStreamPreparer(ctx, faceListID, imageParameter, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceFromStreamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromStream", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceFromStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromStream", resp, "Failure responding to request")
	}

	return
}

// AddFaceFromStreamPreparer prepares the AddFaceFromStream request.
func (client ListClient) AddFaceFromStreamPreparer(ctx context.Context, faceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if targetFace != nil && len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces", pathParameters),
		autorest.WithFile(imageParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddFaceFromStreamSender sends the AddFaceFromStream request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) AddFaceFromStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddFaceFromStreamResponder handles the response to the AddFaceFromStream request. The method always
// closes the http.Response Body.
func (client ListClient) AddFaceFromStreamResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddFaceFromURL add a face to a face list. The input face is specified as an image with a targetFace rectangle. It
// returns a persistedFaceId representing the added face, and persistedFaceId will not expire.
// Parameters:
// faceListID - id referencing a particular face list.
// imageURL - a JSON document with a URL pointing to the image that is to be analyzed.
// userData - user-specified data about the face for any purpose. The maximum length is 1KB.
// targetFace - a face rectangle to specify the target face to be added to a person in the format of
// "targetFace=left,top,width,height". E.g. "targetFace=10,10,100,100". If there is more than one face in the
// image, targetFace is required to specify which face to add. No targetFace means there is only one face
// detected in the entire image.
func (client ListClient) AddFaceFromURL(ctx context.Context, faceListID string, imageURL ImageURL, userData string, targetFace []int32) (result PersistedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: userData,
			Constraints: []validation.Constraint{{Target: "userData", Name: validation.MaxLength, Rule: 1024, Chain: nil}}},
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "AddFaceFromURL", err.Error())
	}

	req, err := client.AddFaceFromURLPreparer(ctx, faceListID, imageURL, userData, targetFace)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddFaceFromURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromURL", resp, "Failure sending request")
		return
	}

	result, err = client.AddFaceFromURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "AddFaceFromURL", resp, "Failure responding to request")
	}

	return
}

// AddFaceFromURLPreparer prepares the AddFaceFromURL request.
func (client ListClient) AddFaceFromURLPreparer(ctx context.Context, faceListID string, imageURL ImageURL, userData string, targetFace []int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	queryParameters := map[string]interface{}{}
	if len(userData) > 0 {
		queryParameters["userData"] = autorest.Encode("query", userData)
	}
	if targetFace != nil && len(targetFace) > 0 {
		queryParameters["targetFace"] = autorest.Encode("query", targetFace, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces", pathParameters),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddFaceFromURLSender sends the AddFaceFromURL request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) AddFaceFromURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddFaceFromURLResponder handles the response to the AddFaceFromURL request. The method always
// closes the http.Response Body.
func (client ListClient) AddFaceFromURLResponder(resp *http.Response) (result PersistedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create create an empty face list. Up to 64 face lists are allowed to exist in one subscription.
// Parameters:
// faceListID - id referencing a particular face list.
// body - request body for creating a face list.
func (client ListClient) Create(ctx context.Context, faceListID string, body NameAndUserDataContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Name", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
				{Target: "body.UserData", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.UserData", Name: validation.MaxLength, Rule: 16384, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, faceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ListClient) CreatePreparer(ctx context.Context, faceListID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ListClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete delete an existing face list according to faceListId. Persisted face images in the face list will also be
// deleted.
// Parameters:
// faceListID - id referencing a particular face list.
func (client ListClient) Delete(ctx context.Context, faceListID string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, faceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ListClient) DeletePreparer(ctx context.Context, faceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ListClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteFace delete an existing face from a face list (given by a persisitedFaceId and a faceListId). Persisted image
// related to the face will also be deleted.
// Parameters:
// faceListID - id referencing a particular face list.
// persistedFaceID - id referencing a particular persistedFaceId of an existing face.
func (client ListClient) DeleteFace(ctx context.Context, faceListID string, persistedFaceID uuid.UUID) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "DeleteFace", err.Error())
	}

	req, err := client.DeleteFacePreparer(ctx, faceListID, persistedFaceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "DeleteFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteFaceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListClient", "DeleteFace", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "DeleteFace", resp, "Failure responding to request")
	}

	return
}

// DeleteFacePreparer prepares the DeleteFace request.
func (client ListClient) DeleteFacePreparer(ctx context.Context, faceListID string, persistedFaceID uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId":      autorest.Encode("path", faceListID),
		"persistedFaceId": autorest.Encode("path", persistedFaceID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}/persistedFaces/{persistedFaceId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteFaceSender sends the DeleteFace request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) DeleteFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteFaceResponder handles the response to the DeleteFace request. The method always
// closes the http.Response Body.
func (client ListClient) DeleteFaceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve a face list's information.
// Parameters:
// faceListID - id referencing a particular face list.
func (client ListClient) Get(ctx context.Context, faceListID string) (result List, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, faceListID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ListClient) GetPreparer(ctx context.Context, faceListID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ListClient) GetResponder(resp *http.Response) (result List, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieve information about all existing face lists. Only faceListId, name and userData will be returned.
func (client ListClient) List(ctx context.Context) (result ListList, err error) {
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.ListClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ListClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/facelists"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ListClient) ListResponder(resp *http.Response) (result ListList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update information of a face list.
// Parameters:
// faceListID - id referencing a particular face list.
// body - request body for updating a face list.
func (client ListClient) Update(ctx context.Context, faceListID string, body NameAndUserDataContract) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: faceListID,
			Constraints: []validation.Constraint{{Target: "faceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "faceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.ListClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, faceListID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "face.ListClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.ListClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ListClient) UpdatePreparer(ctx context.Context, faceListID string, body NameAndUserDataContract) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"faceListId": autorest.Encode("path", faceListID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPathParameters("/facelists/{faceListId}", pathParameters),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ListClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}