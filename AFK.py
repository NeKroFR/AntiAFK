'''
created by NeKro
'''
import time
from pynput.keyboard import Key, Listener, Controller

keyboard = Controller()
delay = 1
i = int('0')

def hold_Z (presskey):
        keyboard.press("z")
        keyboard.release("z")

def hold_S (presskey):
        keyboard.press("s")
        keyboard.release("s")

def hold_Q (presskey):
        keyboard.press("q")

        keyboard.release("q")
z
def hold_D (presskey):    
        keyboard.press("d")
        keyboard.release("d")


while True:
    time.sleep(delay)
    hold_Z(0.5)
    time.sleep(delay)
    hold_S(0.5)
    time.sleep(delay)
    hold_Q(0.5)
    time.sleep(delay)
    hold_D(0.5)
