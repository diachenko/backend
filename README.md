# Calculator service

Simple web service that receives equation and gives response:
```
POST
{
"eq":"222+222"  
}
to receive answer.
```
Each calculation receives unique ID
```
GET returns all calculations in memory.
```

Features:
- + - / * () ^ % "sqrt"  operators support
- sequence of operations is determined by rules. e.g. 2 + 3 * 4 will mutliply 3 by 4 first and then add 2
