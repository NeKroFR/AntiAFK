'''
created by NeKro
'''
import time
from pynput.keyboard import Key, Listener, Controller

keyboard = Controller()
delay = 3


def hold_Z (presskey):
    start = time.time()
    while time.time() - start < presskey:
        keyboard.press("z")

def hold_S (presskey):
    start = time.time()
    while time.time() - start < presskey:
        keyboard.press("s")

def hold_Q (presskey):
    start = time.time()
    while time.time() - start < presskey:
        keyboard.press("q")

def hold_D (presskey):
    start = time.time()
    while time.time() - start < presskey:
        keyboard.press("d")


while True:
    time.sleep(delay)
    hold_Z(0.5)
    time.sleep(delay)
    hold_S(0.5)
    time.sleep(delay)
    hold_Q(0.5)
    time.sleep(delay)
    hold_D(0.5)
