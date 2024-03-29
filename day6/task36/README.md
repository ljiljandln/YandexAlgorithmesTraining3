# 36. Длина кратчайшего пути

||Все языки|Python 3.6|
|---|---|---|
|Ограничение времени|1 секунда|5 секунд|
|Ограничение памяти|64Mb|256Mb|
|Ввод|стандартный ввод или input.txt|   |
|Вывод|стандартный вывод или output.txt|   |

В неориентированном графе требуется найти длину минимального пути между двумя вершинами.

## Формат ввода

Первым на вход поступает число N – количество вершин в графе (1 ≤ N ≤ 100). Затем записана матрица смежности (0 обозначает отсутствие ребра, 1 – наличие ребра). Далее задаются номера двух вершин – начальной и конечной.

## Формат вывода

Выведите L – длину кратчайшего пути (количество ребер, которые нужно пройти).

Если пути нет, нужно вывести -1.

### Пример 1

|Ввод<br><br> ![Скопировать ввод](https://yastatic.net/lego/_/La6qi18Z8LwgnZdsAr1qy1GwCwo.gif)|Вывод<br><br> ![Скопировать вывод](https://yastatic.net/lego/_/La6qi18Z8LwgnZdsAr1qy1GwCwo.gif)|
|---|---|
|10<br>0 1 0 0 0 0 0 0 0 0<br>1 0 0 1 1 0 1 0 0 0<br>0 0 0 0 1 0 0 0 1 0<br>0 1 0 0 0 0 1 0 0 0<br>0 1 1 0 0 0 0 0 0 1<br>0 0 0 0 0 0 1 0 0 1<br>0 1 0 1 0 1 0 0 0 0<br>0 0 0 0 0 0 0 0 1 0<br>0 0 1 0 0 0 0 1 0 0<br>0 0 0 0 1 1 0 0 0 0<br>5 4|2|

### Пример 2

|Ввод<br><br> ![Скопировать ввод](https://yastatic.net/lego/_/La6qi18Z8LwgnZdsAr1qy1GwCwo.gif)|Вывод<br><br> ![Скопировать вывод](https://yastatic.net/lego/_/La6qi18Z8LwgnZdsAr1qy1GwCwo.gif)|
|---|---|
|5<br>0 1 0 0 1<br>1 0 1 0 0<br>0 1 0 0 0<br>0 0 0 0 0<br>1 0 0 0 0<br>3 5|3|
