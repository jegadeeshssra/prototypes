- TESTING CMDs
- v3 commit
    >> curl.exe -i -X GET --header "User-Agent: foobar/1.2.3" http://localhost:4221/user-agent
    >>  curl.exe -i -X GET http://localhost:4221/user-agent

- v4 commit
Start-Job -ScriptBlock { curl -v http://127.0.0.1:4221 }
Start-Job -ScriptBlock { curl -v http://127.0.0.1:4221 }
Start-Job -ScriptBlock { curl -v http://127.0.0.1:4221 }



- v7 commit
curl.exe -i --http1.1  http://localhost:4221/echo/banana 
    \ --next http://localhost:4221/user-agent -H "User-Agent: blueberry/apple-blueberry" 
    \ --next http://localhost:4221/ -H "Connection: close"