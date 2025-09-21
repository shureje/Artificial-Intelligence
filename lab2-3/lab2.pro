:- initialization(main).

parent(pam, bob).
parent(tom, bob).
parent(tom, liz).
parent(bob, ann).
parent(bob, pat).
parent(mary, ann).
parent(pat, juli).

male(tom).
male(bob).
male(jim).
female(mary).
female(liz).
female(pam).
female(pat).
female(ann).
female(juli).

main :-
    repeat,
    write('1 - Все связи'), nl,
    write('2 - Найти родителей'), nl,
    write('3 - Найти внуков'), nl,
    write('4 - Найти матерей'), nl,
    write('5 - Выход'), nl,
    read(Choice),
    process_choice(Choice).

process_choice(1) :- !, show_all, nl, main.
process_choice(2) :- !, find_parent, nl, main.
process_choice(3) :- !, find_grandchildren, nl, main.
process_choice(4) :- !, find_mothers, nl, main.
process_choice(5) :- !, write('Выход').
process_choice(_) :- write('Неверный выбор'), nl, main.

show_all :-
    parent(X,Y),
    write(X), write(' -> '), write(Y), nl, fail.
show_all.

find_parent :-
    write('Введите имя: '),
    read(Name), 
    parent(Parent, Name),
    write('Родитель: '), write(Parent), nl, fail.
find_parent.

find_grandchildren :-
    write('Введите имя: '),
    read(Name),
    parent(Name,Y), parent(Y,X),
    write('Внук: '), write(X), nl, fail.
find_grandchildren.

find_mothers :-
    parent(X,Y), female(X),
    write(X), write(' мать '), write(Y), nl, fail.
find_mothers.
