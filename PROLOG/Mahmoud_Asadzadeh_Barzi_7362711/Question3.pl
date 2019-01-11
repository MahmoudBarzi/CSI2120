distance(Lat1,Lon1,Lat2,Lon2,D) :- 
Lat3 is pi * (Lat1 / 180),
Lon3 is pi * (Lon1 / 180),
Lat4 is pi * (Lat2 / 180),
Lon4 is pi * (Lon2 / 180),
DR is 2 * asin(sqrt(((sin((Lat3-Lat4)/2))^2)+cos(Lat3)*cos(Lat4)*((sin(Lon3-Lon4)/2))^2)),
D is DR * 6371.
