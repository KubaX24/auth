# Auth server
 
Auth server written in go

> Test project

## How it works
- based on jwt tokens to authorized witch are stored in user browser
- when user log in, firstly is created time_token (only for one use), from where user can get jwt token
- database is MongoDB

## File
### _tokens.csv
Format:
``` csv
token_name*,token*,description[optimal]
```
Example:
``` csv
cloudflare_secret,1x0000000000000000000000000000000AA,Turnstile secret token fot testing
```

### _website.csv
Format:
``` csv
website_domain*,token*,after_login_redirect*
```
Example:
``` csv
auth.example.com,very_secret_and_long_token_like_k39nlkk0BNPVmrMTrdBJ_but_longer,auth/loginin.html
auth2.example.com,T9xbhVoZ0cr0H4kypOgbw1iqmft2IemDW99nfdyCiJijCiSKnt,login
```

## Docs
- https://developers.cloudflare.com/turnstile/
- https://jwt.io/introduction

---

Made with love

© KubaX24
