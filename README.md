# Armabot
A simple Discord GOLANG Bot for Arma

1 - create your webhook
https://support.discordapp.com/hc/fr/articles/228383668-Utiliser-les-Webhooks

2 - copy the @armabot directory into your arma3 directory (rename .dll to .so if you use linux)

3 - add this sqf code in your mission
"armabot" callExtension "yourdiscordwebhoookurl;the message to print;yourbotname";