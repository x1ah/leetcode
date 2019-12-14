from queue import Queue


class Foo:
    def __init__(self):
        self.first_queue = Queue(1)
        self.second_queue = Queue(1)
        pass


    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.first_queue.put(True)

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.first_queue.get()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.second_queue.put(True)


    def third(self, printThird: 'Callable[[], None]') -> None:
        self.second_queue.get()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()
