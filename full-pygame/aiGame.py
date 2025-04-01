import tkinter as tk
from tkinter import messagebox
import customtkinter as ctk
from enum import Enum
import random
import time

MIN_START = 40000
MAX_START = 50001


class Number:
    """Хранит число (value), очки (points), банк (bank), финальный счёт (totalPoints) и победителя (winPlayer)."""
    def __init__(self, value, points=0, bank=0):
        self.value = value
        self.points = points
        self.bank = bank
        self.totalPoints = 0
        self.winPlayer = 0

    def calculate_points_and_bank(self):
        x = self.value
        if x % 2 == 0:
            self.points += 1
        else:
            self.points -= 1
        if x % 10 == 0 or x % 5 == 0:
            self.bank += 1

    def clone(self):
        c = Number(self.value, self.points, self.bank)
        c.totalPoints = self.totalPoints
        c.winPlayer = self.winPlayer
        return c


class GameState:

    """Хранит список Number, дочерние состояния (nextStates) и признак final."""
    def __init__(self, current=None):
        self.current = current if current else []
        self.nextStates = []
        self.final = False

    def find_next_state(self):
        any_next = False
        for currentNumber in self.current:
            for val in self.calculate_next_numbers(currentNumber.value):
                any_next = True
                new_num = currentNumber.clone()
                new_num.value = val
                new_num.calculate_points_and_bank()
                child_state = GameState([new_num])
                self.nextStates.append(child_state)
                child_state.find_next_state()
        if not any_next:
            self.final = True



    @staticmethod
    def calculate_next_numbers(x):
        out = []
        if x % 3 == 0:
            out.append(x // 3)
        if x % 4 == 0:
            out.append(x // 4)
        if x % 5 == 0:
            out.append(x // 5)
        return out

    def calculate_game_end(self):
        for nr in self.current:
            if nr.points % 2 == 0:
                nr.totalPoints = nr.points - nr.bank
                nr.winPlayer = 1 if nr.totalPoints % 2 == 0 else 2
            else:
                nr.totalPoints = nr.points + nr.bank
                nr.winPlayer = 1 if nr.totalPoints % 2 == 0 else 2

    def print_endpoints(self):
        """Новый метод из первого кода – выводит конечные узлы дерева."""
        if self.final:
            self.calculate_game_end()
            for nr in self.current:
                print("------------------------------")
                print(f"Value: {nr.value}")
                print(f"Points: {nr.points}")
                print(f"Bank: {nr.bank}")
                print(f"Total Points: {nr.totalPoints}")
                print(f"Winner: {nr.winPlayer} (1=первый игрок, 2=второй игрок)")
                print("------------------------------\n")
        for child in self.nextStates:
            child.print_endpoints()


class Player(Enum):
    USER = 0
    COMPUTER = 1


class Game:
    """Основная логика, хранит GameState и даёт свойства current_number, total_points, bank, final_score."""
    def __init__(self, target_number=7):
        self.moves_history = []
        self.visited_nodes = 0
        self.move_times = []
        self.current_player = Player.USER

        start_num = Number(target_number, 0, 0)
        self.gameState = GameState([start_num])
        self.starting_numbers = self.generate_starting_numbers()

    @property
    def current_number(self):
        if not self.gameState.current:
            return None
        return self.gameState.current[0].value

    @current_number.setter
    def current_number(self, val):
        if self.gameState.current:
            self.gameState.current[0].value = val

    @property
    def total_points(self):
        if not self.gameState.current:
            return 0
        return self.gameState.current[0].points

    @property
    def bank(self):
        if not self.gameState.current:
            return 0
        return self.gameState.current[0].bank

    @property
    def final_score(self):
        if not self.gameState.current:
            return 0
        return self.gameState.current[0].totalPoints

    def set_starting_number(self, number):
        """Подключаем логику из первого кода: генерируем дерево и печатаем концевые узлы."""
        self.gameState = GameState([Number(number, 0, 0)])
        self.gameState.find_next_state()
        # Для наглядности можно сразу посмотреть endpoints:
        self.gameState.print_endpoints()

    def is_divisible_by_345(self, n):
        return n % 3 == 0 and n % 4 == 0 and n % 5 == 0

    def generate_starting_numbers(self):
        valid_nums = []
        for i in range(MIN_START, MAX_START):
            if self.is_divisible_by_345(i):
                valid_nums.append(i)
        return random.sample(valid_nums, 5)

    def is_divisible(self, divisor):
        if not self.gameState.current:
            return False
        return self.gameState.current[0].value % divisor == 0

    def make_move(self, divisor):
        if not self.is_divisible(divisor):
            raise ValueError("Nepareizs gājiens")
        self.moves_history.append((self.current_player, divisor))
        v = self.gameState.current[0].value
        self.gameState.current[0].value = v // divisor
        self.gameState.current[0].calculate_points_and_bank()
        self.switch_player()

    def switch_player(self):
        self.current_player = Player.COMPUTER if self.current_player == Player.USER else Player.USER

    def is_game_over(self):
        for d in [3, 4, 5]:
            if self.is_divisible(d):
                return False
        return True

    def calculate_final_score(self):
        self.gameState.calculate_game_end()

    def evaluate_heuristic(self):
        if not self.gameState.current:
            return 0
        num = self.gameState.current[0]
        score = num.points
        bank = num.bank
        value = num.value
        if value % 2 == 0:
            score += 1
        else:
            score -= 1
        if value % 10 == 0 or value % 10 == 5:
            bank += 1
        if score % 2 == 0:
            score -= bank
        else:
            score += bank
        return score

    def minimax(self, depth, maximizing_player):
        if depth == 0 or self.is_game_over():
            return self.evaluate_heuristic()
        pm = [3, 4, 5]
        if maximizing_player:
            best_score = float('-inf')
            for d in pm:
                if self.is_divisible(d):
                    sv_val = self.gameState.current[0].value
                    sv_pts = self.gameState.current[0].points
                    sv_bnk = self.gameState.current[0].bank
                    self.gameState.current[0].value //= d
                    self.gameState.current[0].calculate_points_and_bank()
                    score_eval = self.minimax(depth - 1, False)
                    self.gameState.current[0].value = sv_val
                    self.gameState.current[0].points = sv_pts
                    self.gameState.current[0].bank = sv_bnk
                    if score_eval > best_score:
                        best_score = score_eval
                    self.visited_nodes += 1
            return best_score
        else:
            worst_score = float('inf')
            for d in pm:
                if self.is_divisible(d):
                    sv_val = self.gameState.current[0].value
                    sv_pts = self.gameState.current[0].points
                    sv_bnk = self.gameState.current[0].bank
                    self.gameState.current[0].value //= d
                    self.gameState.current[0].calculate_points_and_bank()
                    score_eval = self.minimax(depth - 1, True)
                    self.gameState.current[0].value = sv_val
                    self.gameState.current[0].points = sv_pts
                    self.gameState.current[0].bank = sv_bnk
                    if score_eval < worst_score:
                        worst_score = score_eval
                    self.visited_nodes += 1
            return worst_score

    def alphabeta(self, depth, alpha, beta, maximizing_player):
        if depth == 0 or self.is_game_over():
            return self.evaluate_heuristic()
        pm = [3, 4, 5]
        if maximizing_player:
            value = float('-inf')
            for d in pm:
                if self.is_divisible(d):
                    sv_val = self.gameState.current[0].value
                    sv_pts = self.gameState.current[0].points
                    sv_bnk = self.gameState.current[0].bank
                    self.gameState.current[0].value //= d
                    self.gameState.current[0].calculate_points_and_bank()
                    result = self.alphabeta(depth - 1, alpha, beta, False)
                    self.gameState.current[0].value = sv_val
                    self.gameState.current[0].points = sv_pts
                    self.gameState.current[0].bank = sv_bnk
                    if result > value:
                        value = result
                    if value > alpha:
                        alpha = value
                    if beta <= alpha:
                        break
                    self.visited_nodes += 1
            return value
        else:
            value = float('inf')
            for d in pm:
                if self.is_divisible(d):
                    sv_val = self.gameState.current[0].value
                    sv_pts = self.gameState.current[0].points
                    sv_bnk = self.gameState.current[0].bank
                    self.gameState.current[0].value //= d
                    self.gameState.current[0].calculate_points_and_bank()
                    result = self.alphabeta(depth - 1, alpha, beta, True)
                    self.gameState.current[0].value = sv_val
                    self.gameState.current[0].points = sv_pts
                    self.gameState.current[0].bank = sv_bnk
                    if result < value:
                        value = result
                    if value < beta:
                        beta = value
                    if beta <= alpha:
                        break
                    self.visited_nodes += 1
            return value


experiment_results = {
    'computer_wins': 0,
    'human_wins': 0,
    'total_visited_nodes': [],
    'average_move_time': []
}


class GameGUI:
    def __init__(self, game):
        self.game = game
        self.window = ctk.CTk()
        self.window.title("AI game")
        self.window.geometry("1000x600")
        self.window.lift()
        self.custom_font = ctk.CTkFont(family="Arial", size=15)
        self.selected_number_button = None
        self.selected_number = None
        self.toggle_layer = 0
        self.create_del_button()
        self.choose_number()
        self.window.mainloop()

    def create_del_button(self):
        self.dButtonFrame = ctk.CTkFrame(self.window, fg_color='transparent')
        self.dButtonFrame.pack(fill=ctk.BOTH, padx=15, pady=5)
        self.dButton = ctk.CTkButton(self.dButtonFrame, text="Restartēt spēli",
                                     command=self.switch_frames, state=ctk.DISABLED)
        self.dButton.pack(side=tk.RIGHT, padx=15)

    def create_widgets(self):
        self.supermainframe = ctk.CTkFrame(self.window, fg_color='transparent')
        self.supermainframe.pack(pady=10, fill=ctk.BOTH, padx=15)

        self.superframedivider = ctk.CTkFrame(self.supermainframe)
        self.superframedivider.pack(pady=10, padx=15)

        self.score_frame = ctk.CTkFrame(self.superframedivider, fg_color='transparent', width=400)
        self.widgets_frame = ctk.CTkFrame(self.superframedivider, fg_color='transparent', width=400)
        self.history_frame = ctk.CTkFrame(self.superframedivider, fg_color='transparent', width=400)
        self.score_frame.pack(pady=10, padx=15, side=tk.RIGHT)
        self.widgets_frame.pack(pady=10, padx=15, side=tk.RIGHT)
        self.history_frame.pack(pady=10, padx=15, side=tk.RIGHT)

        self.buttons_frame = ctk.CTkFrame(self.widgets_frame, width=500)
        self.players_frame = ctk.CTkFrame(self.widgets_frame)
        self.players_frame.pack(pady=5)
        self.playersname_frame = ctk.CTkFrame(self.widgets_frame)
        self.playersname_frame.pack(pady=5)

        self.first_player_label = ctk.CTkLabel(self.players_frame,
                                               text="Pirmais spēlētajs",
                                               font=ctk.CTkFont("Arial", size=15, weight="bold"),
                                               text_color="#4763c3")
        self.first_player_label.pack(side="left", padx=10)
        self.second_player_label = ctk.CTkLabel(self.players_frame,
                                                text="Otrais spēlētajs",
                                                font=ctk.CTkFont("Arial", size=15, weight="bold"),
                                                text_color="#c35947")
        self.second_player_label.pack(side="right", padx=10)

        self.label = ctk.CTkLabel(self.buttons_frame,
                                  font=ctk.CTkFont("Arial", size=14),
                                  text=f"Spēles skaitlis: {self.game.current_number}")
        self.label.pack(padx=19, pady=10)
        self.buttons_frame.pack()

        self.player_label = ctk.CTkLabel(self.playersname_frame, text="Cilveks        Dators",
                                         font=ctk.CTkFont("Arial", size=15))
        self.player_label.pack(padx=25)

        self.button3 = ctk.CTkButton(self.buttons_frame, text="Dalīt ar 3",
                                     command=lambda: self.on_user_move(3))
        self.button3.pack(pady=5)
        self.button4 = ctk.CTkButton(self.buttons_frame, text="Dalīt ar 4",
                                     command=lambda: self.on_user_move(4))
        self.button4.pack(pady=5)
        self.button5 = ctk.CTkButton(self.buttons_frame, text="Dalīt ar 5",
                                     command=lambda: self.on_user_move(5))
        self.button5.pack(pady=5)

        self.history_label = ctk.CTkLabel(self.history_frame, text="Gājienu vēsture:",
                                          font=ctk.CTkFont("Arial", size=15))
        self.history_label.pack(pady=5, padx=5)

        self.history_text = ctk.CTkTextbox(self.history_frame, width=200, height=250)
        self.history_text.configure(font=("Arial", 17), text_color="green",
                                    border_color="gray", border_width=1, corner_radius=10)
        self.history_text.pack(padx=20, pady=10)

    def create_score_labels(self):
        self.point_frame = ctk.CTkFrame(self.score_frame)
        self.point_frame.configure(fg_color=('white', '#111111'), border_width=1, border_color="gray")
        self.point_frame.pack(pady=25)

        self.total_points_label = ctk.CTkLabel(self.point_frame,
                                               text=f"Kopējie punkti: {self.game.total_points}")
        self.total_points_label.configure(font=ctk.CTkFont("Arial", size=20, weight="bold"),
                                          text_color="#aeb6ff")
        self.total_points_label.pack(pady=5)

        self.plusminus_label = ctk.CTkLabel(self.point_frame, text="+/-")
        self.plusminus_label.configure(font=ctk.CTkFont("Arial", size=20, weight="bold"),
                                       text_color="#aeb6ff")
        self.plusminus_label.pack(pady=5)

        self.bank_points_label = ctk.CTkLabel(self.point_frame,
                                              text=f"Bankas punkti: {self.game.bank}")
        self.bank_points_label.configure(font=ctk.CTkFont("Arial", size=20, weight="bold"),
                                         text_color="#aeb6ff")
        self.bank_points_label.pack(pady=5)

        self.total_points_divider_label = ctk.CTkLabel(self.point_frame, text="                                  ")
        self.total_points_divider_label.configure(font=ctk.CTkFont("Arial", size=20, weight="bold",
                                                                   underline=True, overstrike=True))
        self.total_points_divider_label.pack(padx=20, pady=2)

        self.final_points_label = ctk.CTkLabel(self.point_frame,
                                               text=f"Gala rezultats: {self.game.final_score}")
        self.final_points_label.configure(font=ctk.CTkFont("Arial", size=20, weight="bold"),
                                          text_color="#aeb6ff")
        self.final_points_label.pack(pady=7)

    def update_score_labels(self):
        tp_text = f"Kopējie punkti: {self.game.total_points}"
        tp_color = "#4763c3" if self.game.total_points % 2 == 0 else "#c35947"
        self.total_points_label.configure(text=tp_text, text_color=tp_color)
        self.bank_points_label.configure(text=f"Bankas punkti: {self.game.bank}")

    def update_labels_and_buttons(self):
        self.label.configure(text=f"Pašreizējais skaitlis: {self.game.current_number}")
        self.update_buttons()
        self.update_score_labels()

    def update_buttons(self):
        for button, number in [(self.button3, 3), (self.button4, 4), (self.button5, 5)]:
            if self.game.is_divisible(number):
                button.configure(fg_color="#50C878", text_color="black", state="normal")
            else:
                button.configure(fg_color="#DC143C", text_color="black", state="disabled")

    def choose_number(self):
        try:
            self.ultramainframe.destroy()
        except:
            pass

        self.ultramainframe = ctk.CTkFrame(self.window, fg_color='transparent')
        self.ultramainframe.pack(pady=10, fill=tk.BOTH, padx=15)

        self.mainframe = ctk.CTkFrame(self.ultramainframe)
        self.mainframe.pack()

        self.tabview = ctk.CTkTabview(self.mainframe, width=300)
        self.tabview.configure(height=100)
        self.tabview.add("Izvēlēties")
        self.tabview.add("Ievadīt")
        self.tabview.tab("Ievadīt").grid_columnconfigure(0, weight=5)
        self.tabview.tab("Izvēlēties").grid_columnconfigure(0, weight=1)
        self.tabview.pack()

        self.insertnumber = ctk.CTkLabel(self.tabview.tab("Ievadīt"),
                                         text="Ievadiet skaitli (40000 - 50000):")
        self.insertnumber.configure(font=("Arial", 18), text_color="#84888e")
        self.insertnumber.pack()

        number_entry = ctk.CTkEntry(self.tabview.tab("Ievadīt"))
        number_entry.pack(pady=20)

        self.quinbuttonframe = ctk.CTkFrame(self.tabview.tab("Izvēlēties"),
                                            fg_color='transparent',
                                            height=40)
        self.quinbuttonframe.pack(pady=10, fill=tk.BOTH)

        choose_button_style = {
            'text_color': 'white',
            'font': self.custom_font,
            'fg_color': '#3061bf',
            'hover_color': '#659bde',
            'width': 10,
            'border_width': 1,
            'border_color': "#547ccb"
        }

        def create_button(num):
            button = ctk.CTkButton(self.quinbuttonframe, text=str(num))
            button.configure(**choose_button_style, command=lambda n=num, btn=button: select_number(n, btn))
            button.pack(pady=1, side=tk.RIGHT, padx=2)
            return button

        def select_number(num, btn):
            if self.selected_number_button:
                self.selected_number_button.configure(text_color="white", fg_color="#3061bf", hover_color="#31559b")
            self.selected_number = num
            self.selected_number_button = btn
            btn.configure(text_color="black", fg_color="#50C879", hover_color="#079838")
            update_cancel_button()

        def cancel_choice():
            if self.selected_number_button:
                self.selected_number = None
                update_cancel_button()
                reset_button_colors()

        def update_cancel_button():
            if self.selected_number is not None:
                cancel_button.pack(pady=10)
            else:
                cancel_button.pack_forget()

        def reset_button_colors():
            for b in button_list:
                b.configure(text_color="white", fg_color="#3061bf", hover_color="#31559b")

        cancel_button = ctk.CTkButton(self.tabview.tab("Izvēlēties"), text="Atcelt izvēli",
                                      command=cancel_choice, **choose_button_style)
        self.textstartchoose = ctk.CTkLabel(self.tabview.tab("Izvēlēties"), text="Izvēlies sākuma skaitli")
        self.textstartchoose.configure(font=("Arial", 18))
        self.textstartchoose.pack()

        button_list = []
        for val in self.game.starting_numbers:
            b = create_button(val)
            button_list.append(b)

        ultraframedivider = ctk.CTkFrame(self.mainframe)
        ultraframedivider.pack(pady=10, padx=15)

        self.starter_frame = ctk.CTkFrame(ultraframedivider, fg_color='transparent')
        self.algorythm_frame = ctk.CTkFrame(ultraframedivider, fg_color='transparent')
        self.starter_frame.pack(pady=10, padx=10, side=tk.RIGHT)
        self.algorythm_frame.pack(pady=10, padx=10, side=tk.RIGHT)

        self.whoplays = ctk.CTkLabel(self.starter_frame, text="Kurš sāk spēli:")
        self.whoplays.configure(font=("Arial", 18))
        self.whoplays.pack(pady=5)

        starter_var = tk.IntVar(value=0)

        def on_submit():
            try:
                if self.selected_number is not None and number_entry.get():
                    messagebox.showerror("Kļūda",
                        "Jums ir jāizvēlās ģenerēts vai pašam ievadīts skaitlis!")
                    self.reset_game()
                elif self.selected_number is not None:
                    number = self.selected_number
                else:
                    number = int(number_entry.get())

                if 40000 <= number <= 50000:
                    if number % 3 == 0 and number % 4 == 0 and number % 5 == 0:
                        self.game.set_starting_number(number)
                        self.game.current_player = Player(starter_var.get())
                        self.switch_frames()
                        self.update_labels_and_buttons()
                        if self.game.current_player == Player.COMPUTER:
                            self.computer_move()
                    else:
                        messagebox.showerror("Kļūda",
                            "Skaitlim jābūt dalāmam ar 3, 4 un 5 bez atlikuma!")
                        self.choose_number()
                else:
                    messagebox.showerror("Kļūda",
                        "Skaitlim jābūt diapazonā no 40000 līdz 50000!")
                    self.choose_number()
            except ValueError:
                messagebox.showerror("Kļūda", "Ievadiet derīgu veselu skaitli!")
                self.choose_number()

        def set_pl_labels_positions():
            if starter_var.get() == 0:
                self.player_label.configure(text="Cilveks        Dators")
            else:
                self.player_label.configure(text="Dators        Cilveks")

        self.gamer = ctk.CTkRadioButton(self.starter_frame, text="Cilveks",
                                        variable=starter_var,
                                        value=0,
                                        command=set_pl_labels_positions)
        self.gamer.configure(font=("Arial", 16))
        self.gamer.pack(pady=5)

        self.pc = ctk.CTkRadioButton(self.starter_frame, text="Dators",
                                     variable=starter_var,
                                     value=1,
                                     command=set_pl_labels_positions)
        self.pc.configure(font=("Arial", 16))
        self.pc.pack(pady=5)

        self.choosealgorithm = ctk.CTkLabel(self.algorythm_frame, text="Izvēlieties algoritmu:")
        self.choosealgorithm.configure(font=("Arial", 18))
        self.choosealgorithm.pack(pady=5)

        self.algorithm_var = tk.StringVar(value="minimax")
        self.mini = ctk.CTkRadioButton(self.algorythm_frame, text="Minimax",
                                       variable=self.algorithm_var,
                                       value="minimax")
        self.mini.configure(font=("Arial", 16))
        self.mini.pack(pady=5)

        self.alfa = ctk.CTkRadioButton(self.algorythm_frame, text="Alpha-Beta",
                                       variable=self.algorithm_var,
                                       value="alphabeta")
        self.alfa.configure(font=("Arial", 16))
        self.alfa.pack(pady=5)

        start_button = ctk.CTkButton(self.mainframe,
                                     text="Sākt spēli",
                                     command=on_submit,
                                     **choose_button_style)
        start_button.configure(width=40, height=45)
        start_button.pack(pady=15)

    def show_final_message(self, final_message):
        result_window = ctk.CTkToplevel(self.window)
        self.window.title("AI game")
        ctk.CTkLabel(result_window, text=final_message).pack(pady=10)

        def restart_game():
            result_window.destroy()
            self.reset_game()

        def close_result_window():
            result_window.destroy()

        tk.Button(result_window, text="Sākt vēlreiz", command=restart_game).pack(side=tk.LEFT, padx=10, pady=10)
        tk.Button(result_window, text="Pabeigt", command=close_result_window).pack(side=tk.RIGHT, padx=10, pady=10)

    def add_end_game_buttons(self):
        self.restart_button = tk.Button(self.window, text="Sākt vēlreiz", command=self.reset_game)
        self.restart_button.pack(pady=5)

        def close_program():
            self.window.destroy()

        self.exit_button = tk.Button(self.window, text="Pabeigt", command=close_program)
        self.exit_button.pack(pady=5)

    def reset_game(self):
        self.game = Game()
        self.selected_number = None
        self.selected_number_button = None

    def on_user_move(self, number):
        try:
            self.game.make_move(number)
            self.update_labels_and_buttons()
            self.update_history()
            if not self.game.is_game_over():
                self.computer_move()
            else:
                self.game.calculate_final_score()
                if self.game.total_points % 2 == 0:
                    self.plusminus_label.configure(text="-")
                else:
                    self.plusminus_label.configure(text="+")
                gr_text = f"Gala rezultats: {self.game.final_score}"
                gr_color = "#4763c3" if self.game.final_score % 2 == 0 else "#c35947"
                self.final_points_label.configure(text=gr_text, text_color=gr_color)
                final_message = "Jūs uzvarējāt!" if self.game.final_score % 2 == 0 else "Dators uzvarēja!"
                self.show_final_results(final_message)
                self.reset_game()
        except ValueError:
            messagebox.showerror("Kļūda", "Nederīgs gājiens")

    def computer_move(self):
        start_time = time.perf_counter()
        best_move = None
        best_eval = float('-inf')
        for num in [3, 4, 5]:
            if self.game.is_divisible(num):
                original_value = self.game.current_number
                self.game.current_number //= num
                self.game.gameState.current[0].calculate_points_and_bank()
                if self.algorithm_var.get() == "minimax":
                    eval_ = self.game.minimax(3, False)
                else:
                    eval_ = self.game.alphabeta(3, float('-inf'), float('inf'), False)
                self.game.current_number = original_value
                if eval_ > best_eval:
                    best_eval = eval_
                    best_move = num
        if best_move is not None:
            self.game.make_move(best_move)
            self.update_labels_and_buttons()
            self.update_history()
            eval_time = time.perf_counter() - start_time
            self.game.move_times.append(eval_time)
            if self.game.is_game_over():
                self.game.calculate_final_score()
                if self.game.total_points % 2 == 0:
                    self.plusminus_label.configure(text="-")
                else:
                    self.plusminus_label.configure(text="+")
                gr_text = f"Gala rezultats: {self.game.final_score}"
                gr_color = "#47c5c5" if self.game.final_score % 2 == 0 else "#9b2a0a"
                self.final_points_label.configure(text=gr_text, text_color=gr_color)
                final_message = "Dators uzvarēja!" if self.game.final_score % 2 == 0 else "Jūs uzvarējāt!"
                self.show_final_results(final_message)
                self.reset_game()
        else:
            messagebox.showerror("Kļūda", "Dators nevarēja veikt gājienu")

    def show_final_results(self, winner_message):
        print("Move times:", self.game.move_times)
        avg_time = sum(self.game.move_times) / len(self.game.move_times) if self.game.move_times else 0
        visited_nodes = self.game.visited_nodes
        additional_message = f"\nDatora vidējais laiks gājienu izpildei: {avg_time:.2f} sek.\nApmeklēto virsotņu skaits: {visited_nodes}"
        final_message = winner_message + additional_message
        messagebox.showinfo("Spēle beigusies", final_message)
        self.reset_game()

    def update_history(self):
        self.history_text.delete(1.0, tk.END)
        if not self.game.moves_history:
            return
        first_player = self.game.moves_history[0][0]
        if first_player == Player.USER:
            user_color = "#4763c3"
            computer_color = "#c35947"
        else:
            user_color = "#c35947"
            computer_color = "#4763c3"
        for pl, mv in self.game.moves_history:
            player_name = "Lietotājs" if pl == Player.USER else "Dators"
            color = user_color if pl == Player.USER else computer_color
            self.history_text.insert(tk.END, f"{player_name} ", f"{player_name}_color")
            self.history_text.insert(tk.END, f"dalīja ar {mv}\n")
            self.history_text.tag_config(f"{player_name}_color", foreground=color)

    def switch_frames(self):
        if self.toggle_layer == 1:
            self.dButton.configure(state=ctk.DISABLED)
            self.toggle_layer = 0
            self.supermainframe.destroy()
            self.choose_number()
            self.reset_game()
        else:
            self.ultramainframe.destroy()
            self.dButtonFrame.destroy()
            self.create_del_button()
            self.dButton.configure(state=ctk.NORMAL)
            self.create_widgets()
            self.create_score_labels()
            self.toggle_layer = 1


if __name__ == "__main__":
    g = Game()
    gui = GameGUI(g)
