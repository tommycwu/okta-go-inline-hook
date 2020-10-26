# okta-go-inline-hook

This is a sample okta inline saml hook written in golang as a aws lambda function.  https://developer.okta.com/docs/concepts/inline-hooks/.  

1) You will need an okta org.  https://developer.okta.com/signup/

2) Next set up a an inline hook in your org. https://help.okta.com/en/prod/Content/Topics/automation-hooks/add-inline-hooks.htm

3) The inline hook will be triggered by a saml assertion from an app in your org https://developer.okta.com/docs/guides/build-sso-integration/saml2/create-your-app/. 

4) The inline hook will reach out to your REST based API (webservice) with an endpoint of /samlhook.

5) The sample app needs to be hosted in a publicly accessible address with an ssl/https enabled port 
  - This sample includes a makefile and toml file set up specifically for https://www.netlify.com/.

6) In your API, it will need to take an incoming request and in the return body you will need to send it a json string that the inline hook will recognize (https://developer.okta.com/docs/reference/saml-hook/).

7) You can test this with samltracer (https://chrome.google.com/webstore/detail/saml-tracer/mpdajninpobndbfcldcmbpnnbhibjmch?hl=en, https://addons.mozilla.org/en-US/firefox/addon/saml-tracer/)
