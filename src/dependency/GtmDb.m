readChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do readChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set json=json_"}"
 ...write json
 ..else  do
 ...set json=json_","
 ...do readChain(x)
 ...quit
 .else  do
 ..do readChain(string)
 .quit
 quit

read
 set recordCount=0
 set json="{"
 set s=$$firstDescendant()
 if s="""""" do
 .write "wrong database name"
 else  do
 .do readChain(s)
 .quit
 quit

limitChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do limitChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set json=json_"}"
 ...write json
 ..else  if recordCount<recordLimit do
 ...set json=json_","
 ...do limitChain(x)
 ..else  do
 ...set json=json_"}"
 ...write json
 ...quit
 .else  do
 ..do limitChain(string)
 .quit
 quit

limit(lim)
 set recordCount=0
 set recordLimit=lim
 set json="{"
 set s=$$firstDescendant()
 if s="""""" do
 .write "wrong database name"
 else  do
 .do limitChain(s)
 .quit
 quit

firstChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do firstChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set json=json_"}"
 ...write json
 ..else  do
 ...set json=json_","
 ...if $piece(x,",")=$$firstDescendant() do
 ....do firstChain(x)
 ...else  do
 ....set json=$extract(json,1,$length(json)-1)
 ....set json=json_"}"
 ....write !,json,!
 ....quit
 ...quit
 .else  do
 ..do firstChain(string)
 .quit
 quit

first
 set recordCount=0
 set recordLimit=1
 set json="{"
 set s=$$firstDescendant()
 if s="""""" do
 .write "wrong database name"
 else  do
 .do firstChain(s)
 .quit
 quit

lastChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do lastChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set json=json_"}"
 ...write json
 ..else  do
 ...set json=json_","
 ...if $piece(x,",")=""""_$$lastDescendant("")_"""" do
 ....do lastChain(x)
 ...else  do
 ....set json=$extract(json,1,$length(json)-1)
 ....set json=json_"}"
 ....write !,json,!
 ....quit
 ...quit
 .else  do
 ..do lastChain(string)
 .quit
 quit

last
 set recordCount=0
 set recordLimit=1
 set json="{"
 set s=$$lastDescendant("")
 if s="" do
 .write "wrong database name"
 else  do
 .do lastChain(""""_s_"""")
 .quit
 quit

byIdChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do byIdChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set json=json_"}"
 ...write json
 ...set resultJson=json
 ..else  if recordCount<recordLimit do
 ...set json=json_","
 ...do byIdChain(x)
 ..else  do
 ...set json=json_"}"
 ...write json
 ...set resultJson=json
 ...quit
 .else  do
 ..do byIdChain(string)
 .quit
 quit

byId(id)
 set recordCount=0
 set recordLimit=1
 set findById=id
 set json="{"
 if id="" do
 .write "Invalid id"
 else  do
 .do byIdChain(""""_id_"""")
 .quit
 quit

delete(id)
 set ret=0
 xecute "kill ^"_^db_"("""_id_""")"
 set ret=1
 quit ret

byIdWithoutPrintChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do byIdWithoutPrintChain(string)
 else  if $$isDirect(string) do
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...set resultJson=json
 ..else  if recordCount<recordLimit do
 ...set json=json_","
 ...do byIdWithoutPrintChain(x)
 ..else  do
 ...set resultJson=json
 ...quit
 .else  do
 ..do byIdWithoutPrintChain(string)
 .quit
 quit

byIdWithoutPrint(id)
 set recordCount=0
 set recordLimit=1
 set findById=id
 set json=""
 if id="" do
 .write "Invalid id"
 else  do
 .do byIdWithoutPrintChain(""""_id_"""")
 .quit
 quit

whereChain(string)
 if $$hasDescendants(string) do
 .set json=$$appendDescendantField(json,$$back(string,","))
 .set string=$$findInitial(string)
 .do whereChain(string)
 else  if $$isDirect(string) do
 .; write !,$$findValue(string)>(""""_$piece(recordQuery," ",3,$l(recordQuery," "))_""""),!
 .set boolv=0
 .if ($piece(recordQuery," ",2,2)="==")!($piece(recordQuery," ",2,2)="=") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($$findValue(string)=(""""_$piece(recordQuery," ",3,$l(recordQuery," "))_""""))
 .else  if ($piece(recordQuery," ",2,2)=">") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($extract($$findValue(string),2,$l($$findValue(string))-1)>($piece(recordQuery," ",3,$l(recordQuery," "))))
 .else  if ($piece(recordQuery," ",2,2)="<") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($extract($$findValue(string),2,$l($$findValue(string))-1)<($piece(recordQuery," ",3,$l(recordQuery," "))))
 .else  if ($piece(recordQuery," ",2,2)="!=") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($extract($$findValue(string),2,$l($$findValue(string))-1)'=($piece(recordQuery," ",3,$l(recordQuery," "))))
 .else  if ($piece(recordQuery," ",2,2)=">=") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($extract($$findValue(string),2,$l($$findValue(string))-1)>=($piece(recordQuery," ",3,$l(recordQuery," "))))
 .else  if ($piece(recordQuery," ",2,2)="<=") do
 ..set boolv=(($$back(string,","))=(""""_$piece(recordQuery," ")_""""))&($extract($$findValue(string),2,$l($$findValue(string))-1)<=($piece(recordQuery," ",3,$l(recordQuery," "))))
 ..quit
 .if boolv do
 ..set s=$piece(string,",")
 ..do byIdWithoutPrint($extract(s,2,$l(s)-1))
 ..set whereJson=whereJson_resultJson_","
 ..quit
 .set json=$$appendDirectField(json,$$back(string,","))
 .set json=$$appendValue(json,$$findValue(string))
 .set string=$$findNext(string)
 .if $$back(string,",")="""""" do
 ..set json=$extract(json,1,$length(json)-1)
 ..set x=$$checkForBack(string)
 ..if x=5 do
 ...if $extract(whereJson,$l(whereJson),$l(whereJson))="," do
 ....set whereJson=$extract(whereJson,1,$length(whereJson)-1)
 ....quit
 ...set whereJson=whereJson_"}}"
 ...write whereJson
 ..else  do
 ...set json=json_","
 ...do whereChain(x)
 ...quit
 .else  do
 ..do whereChain(string)
 .quit
 quit

where(query)
 set recordCount=0
 set recordQuery=query
 set json="{"
 set whereJson="{""matchRec"":{"
 set resultJson=""
 set s=$$firstDescendant()
 if s="""""" do
 .write "wrong database name"
 else  do
 .do whereChain(s)
 .quit
 quit

checkForBack(string)
 if $$front(string,",")="""""" do
 .set string=5
 else  do
 .if $$back(string,",")="""""" do
 ..if $l($$front(string,","),",")=1 do
 ...set recordCount=recordCount+1
 ...quit
 ..set json=json_"}"
 ..set string=$$front(string,",")
 ..set string=$$findNext(string)
 ..set string=$$checkForBack(string)
 ..quit
 .quit
 quit string

front(string,separator)
 set ret=""
 if $l(string,separator)=1 do
 .set ret=$piece(string,separator)
 else  do
 .set ret=$piece(string,separator,0,$l(string,separator)-1)
 quit ret

back(string,separator)
 quit $piece(string,separator,$l(string,separator),$l(string,separator))

hasDescendants(string)
 xecute "set s=$data(^"_^db_"("_string_"))"
 set boolean=0
 if (s=10)!(s=11) do
 .set boolean=1
 .quit
 quit boolean

isDirect(string)
 xecute "set s=$data(^"_^db_"("_string_"))"
 set boolean=0
 if s=1 do
 .set boolean=1
 .quit
 quit boolean

appendDescendantField(store,field)
 quit store_field_":{"

appendDirectField(store,field)
 quit store_field_":"

appendValue(store,value)
 quit store_value_","

findInitial(string)
 xecute "set s=$order(^"_^db_"("_string_",""""))"
 quit string_","""_s_""""

findNext(string)
 set returnValue=""
 xecute "set s=$order(^"_^db_"("_string_"))"
 if $length(string,",")>1 do
 .set returnValue=$$front(string,",")_","""_s_""""
 else  do
 .set returnValue=""""_s_""""
 quit returnValue

findValue(string)
 xecute "set s=^"_^db_"("_string_")"
 quit """"_s_""""

firstDescendant()
 xecute "set string=$order(^"_^db_"(""""))"
 quit """"_string_""""

lastDescendant(s)
 xecute "set string=$order(^"_^db_"("""_s_"""))"
 if string="" do
 .set string=s
 else  do
 .set string=$$lastDescendant(string)
 .quit
 quit string

nextId
 set count=0
 set s=""
 for i=0:1 set count=count+1 xecute "set s=$order(^"_^db_"("""_s_"""))" if s="" quit
 write count
 quit

count
 set count=0
 set s=""
 for i=0:1 set count=count+1 xecute "set s=$order(^"_^db_"("""_s_"""))" if s="" quit
 write count-1
 quit
