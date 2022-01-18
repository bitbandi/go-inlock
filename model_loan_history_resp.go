/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.  # Responses Endpoints return HTTP status code 200 in case of successful operation. The following response body format is used:  {   \"result\": {\"status\": \"ok\", \"`ENDPOINT`\": `RESPONSE`} }  E.g.: {   \"result\": {     \"status\": \"ok\",     \"getAutoLendApr\": {       \"apr_usdc\": 8.5,       \"apr_ilk\": 10.9     }   } }  In case of error HTTP status code 422 is returned. In this case the following response body format applies:  {   \"result\": {\"status\": \"error\", \"`ENDPOINT`\": {}},   \"error\": {\"code\": \"`ERROR_CODE`\", \"message\": \"`ERROR_MESSAGE`\"} }  E.g.: {   \"result\": {     \"status\": \"error\",     \"login\": {}   },   \"error\": {     \"code\": \"EGEN_badreq_E001\",     \"message\": \"Bad request or missing field, please check the API documentation\"   } }  The following general error codes are defined: - `EGEN_dbresp_E001`: Internal error. - `EGEN_badreq_E001`: Missing or invalid parameters passed in request body. - `EGEN_undef_E001`: Unknown error.  In case of authentication error status code 401 is returned with the following body:  {   \"result\": {\"status\": \"error\", \"`ENDPOINT`\": {}},   \"error\": {\"code\": \"EGEN_unauth_E001\", \"message\": \"Unauthorized access\"} }  # Authentication (valid only for retail endpoints) This API offers an APIKey + Signed with Secret Key based authentication.  All gateway API requests must include the proper authentication headers:  Every API call has a SHA-512 HMAC signature generated with partner's secret key.  Gateway backend generates it's own HMAC signature and compares it with the partner.  Unauthorized access status returns if both hash don't match.  The HMAC signature is sent as a HTTP header called 'X-Signature'.  Also mandatory for all API calls is the partner's APIKey in the HTTP header called 'X-Apikey'.  POST data signature remark: When request contains POST data, is should be serialized json without  whitespaces. Signature calculated based on this raw string: concat(full url + sha256(serialized json post data))  Example code to create/check a signature in python:  ``` def reproduce_signature(secret):     raw_data = str(request.url+serialized_data()).encode('utf-8')     raw_secret = base58.b58decode(secret)     raw_check = hmac.new(raw_secret, raw_data, hashlib.sha512)     check = base58.b58encode(raw_check.digest())     return check.decode('utf-8')   def serialized_data():     if not request.is_json:         return ''     # most compact format of a json, without any whitespace     return json.dumps(request.json, indent=None, separators=(',', ':')) ```  <SecurityDefinitions /> 
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type LoanHistoryResp struct {
	// List of event records. 
	Events []LoanHistoryRespEvents `json:"events,omitempty"`
}
