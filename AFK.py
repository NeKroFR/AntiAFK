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
        time.sleep(5)
        keyboard.release("z")

def hold_S (presskey):
    for i in range(250):
        keyboard.press("s")
        time.sleep(5)
        keyboard.release("s")

def hold_Q (presskey):
    for i in range(250):
        keyboard.press("q")
        time.sleep(5)
        keyboard.release("q")
z
def hold_D (presskey):    
    for i in range(250):
        keyboard.press("d")
        time.sleep(5)
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
