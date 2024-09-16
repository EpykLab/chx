Output useChxWithPipes.gif

Require echo

Set Shell "bash"
Set FontSize 32
Set Width 1920
Set Height 1080

Type "eza -liR | awk '{print $8}' | iters -c sha256sum 2> /dev/null | awk '{print $1}' | iters -p -c chx vt | jq -r .data.attributes.last_analysis_results.ALYac.result" Sleep 500ms  Enter

Sleep 5s
