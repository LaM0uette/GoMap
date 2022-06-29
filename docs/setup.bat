chcp 1252>nul
@echo off
cls

echo.
echo Ajout de la variable en cours..
setx GoMap "T:\- 4 Suivi Appuis\26_MACROS\GOGoMap\GoMap.exe"
echo.
start "" "T:\- 4 Suivi Appuis\26_MACROS\GO\GoMap\GoMap_REG.reg"

pause