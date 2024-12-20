/*
 * travel.pl
 * By Antonio F. Huertas
 * This is a knowledge base of fictitious travel information.
 */

% by_car(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by car.
by_car(bayamon, guaynabo).
by_car(guaynabo, san_juan).
by_car(san_juan, carolina).
by_car(aguadilla, mayaguez).
by_car(mayaguez, san_german).

% by_train(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by train.
by_train(bayamon, aguadilla).
by_train(aguadilla, cabo_rojo).
by_train(san_german, ponce).
by_train(carolina, fajardo).
by_train(carolina, ponce).

% by_plane(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by plane.
by_plane(carolina, orlando).
by_plane(carolina, new_york).
by_plane(carolina, paris).
by_plane(paris, tokio).
by_plane(new_york, san_antonio).
by_plane(san_antonio, san_francisco).

% travelByCar(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by car.
travelByCar(Origin, Destination) :-
    by_car(Origin, Destination).
travelByCar(Origin, Destination) :-
    by_car(Origin, Intermediate),
    travelByCar(Intermediate, Destination).

% travelByTrain(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by train.
travelByTrain(Origin, Destination) :-
    by_train(Origin, Destination).
travelByTrain(Origin, Destination) :-
    by_train(Origin, Intermediate),
    travelByTrain(Intermediate, Destination).

% travelByPlane(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by plane.
travelByPlane(Origin, Destination) :-
    by_plane(Origin, Destination).
travelByPlane(Origin, Destination) :-
    by_plane(Origin, Intermediate),
    travelByPlane(Intermediate, Destination).

% travel(?Origin, ?Destination)
% Succeeds if Destination is reachable from Origin by car, train, or plane.
travel(Origin, Destination) :-
    travelByCar(Origin, Destination).
travel(Origin, Destination) :-
    travelByTrain(Origin, Destination).
travel(Origin, Destination) :-
    travelByPlane(Origin, Destination).
travel(Origin, Destination) :-
    (by_car(Origin, Intermediate); by_train(Origin, Intermediate); by_plane(Origin, Intermediate)),
    travel(Intermediate, Destination).