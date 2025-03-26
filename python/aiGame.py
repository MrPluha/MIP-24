import sys
import random

MAX_START = 50000
MIN_START = 40000

class Number:
    def __init__(self, value, points=0, bank=0):
        self.value = value
        self.points = points
        self.bank = bank
        self.total_points = 0
        self.win_player = 0  # 0: no winner, 1: first player wins, 2: second player wins

    def calculate_points_and_bank(self):
        if self.value % 2 == 0:
            self.points += 1
        else:
            self.points -= 1

        if self.value % 10 == 0 or self.value % 5 == 0:
            self.bank += 1

    def calculate_next_numbers(self):
        out = []
        if self.value % 3 == 0:
            out.append(self.value // 3)
        if self.value % 4 == 0:
            out.append(self.value // 4)
        if self.value % 5 == 0:
            out.append(self.value // 5)
        return out

class GameState:
    def __init__(self, numbers):
        self.current = numbers
        self.final = False
        self.next_states = []

    def find_next_state(self):
        for current_number in self.current:
            if current_number is None:
                continue

            next_numbers = []
            next_values = current_number.calculate_next_numbers()
            for v in next_values:
                new_number = Number(v, current_number.points, current_number.bank)
                new_number.calculate_points_and_bank()
                next_numbers.append(new_number)

            if not next_numbers:
                continue

            next_state = GameState(next_numbers)
            self.next_states.append(next_state)
            next_state.find_next_state()

        if not self.next_states:
            self.final = True

    def calculate_game_end(self):
        for nr in self.current:
            if nr.points % 2 == 0:
                nr.total_points = nr.points - nr.bank
                nr.win_player = 1 if nr.total_points % 2 == 0 else 2
            else:
                nr.total_points = nr.points + nr.bank
                nr.win_player = 1 if nr.total_points % 2 == 0 else 2

    def print_endpoints(self):
        if self.final:
            self.calculate_game_end()
            for nr in self.current:
                print(f"Value: {nr.value}\nPoints: {nr.points}; Bank: {nr.bank}; Total Points: {nr.total_points}")
                print(f"Player {nr.win_player} would win (1 = first player, 2 = second player)\n")
        
        for next_state in self.next_states:
            next_state.print_endpoints()


def generate_start_numbers(seed):
    int_in_range = [i for i in range(MIN_START, MAX_START) if i % 3 == 0 and i % 4 == 0 and i % 5 == 0]
    random.seed(seed)
    return random.sample(int_in_range, 5)

def get_start_values(args):
    if args:
        return [int(arg) for arg in args]
    return generate_start_numbers(random.randint(1, 9999999999999999999))

def prepare_game_start(start_numbers):
    numbers = [Number(value) for value in start_numbers]
    return GameState(numbers)

def main():
    start_values = get_start_values(sys.argv[1:])
    start_state = prepare_game_start(start_values)
    
    print("GAME START VALUES")
    for number in start_state.current:
        print(number.value)
    
    start_state.find_next_state()
    start_state.print_endpoints()

if __name__ == "__main__":
    main()
