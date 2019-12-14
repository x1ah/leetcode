from queue import Queue


class FizzBuzz:
    def __init__(self, n: int):
        self.n = n
        self.fizz_continue = Queue(1)
        self.buzz_continue = Queue(1)
        self.fizzbuzz_continue = Queue(1)
        self.number_continue = Queue(1)
        self.number_continue.put(True)
    
    @classmethod
    def is_fizz(cls, n):
        return (n % 3 == 0) and (n % 5 != 0)
    
    @classmethod
    def is_buzz(cls, n):
        return (n % 5 == 0) and (n % 3 != 0)
    
    @classmethod
    def is_fizzbuzz(cls, n):
        return (n % 3 == 0) and (n % 5 == 0)
    
    @classmethod
    def is_number(cls, n):
        return (n % 3 != 0) and (n % 5 != 0)
    
    def rouge(self, you_queue, condition, printer):
        for i in range(1, self.n + 1):
            if condition(i):
                you_queue.get()
                printer(i)
                next_queue = self.get_next_queue(i + 1)
                next_queue.put(True)
                
    def get_next_queue(self, n):
        result_map = {
            self.is_buzz(n): self.buzz_continue,
            self.is_fizz(n): self.fizz_continue,
            self.is_fizzbuzz(n): self.fizzbuzz_continue,
            self.is_number(n): self.number_continue,
        }
        return result_map[True]
    
    # printFizz() outputs "fizz"
    def fizz(self, printFizz: 'Callable[[], None]') -> None:
        self.rouge(self.fizz_continue, self.is_fizz, lambda _: printFizz())
        
    # printBuzz() outputs "buzz"
    def buzz(self, printBuzz: 'Callable[[], None]') -> None:
        self.rouge(self.buzz_continue, self.is_buzz, lambda _: printBuzz())
        
    # printFizzBuzz() outputs "fizzbuzz"
    def fizzbuzz(self, printFizzBuzz: 'Callable[[], None]') -> None:
        self.rouge(self.fizzbuzz_continue, self.is_fizzbuzz, lambda _: printFizzBuzz())

    # printNumber(x) outputs "x", where x is an integer.
    def number(self, printNumber: 'Callable[[int], None]') -> None:
        self.rouge(self.number_continue, self.is_number, lambda i: printNumber(i))

