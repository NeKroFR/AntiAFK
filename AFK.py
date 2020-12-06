'''
created by NeKro
'''
import time
from pynput.keyboard import Key, Listener, Controller


delay = 3
presskey = 3
keyboard = Controller()

while True:
    time.sleep(delay)
    keyboard.press("z")
    time.sleep(presskey)
    keyboard.release("z") 
    time.sleep(delay)
    keyboard.press("s")
    time.sleep(presskey)
    keyboard.release("s") 
    time.sleep(delay)
    keyboard.press("q")
    time.sleep(presskey)
    keyboard.release("q") 
    time.sleep(delay)
    keyboard.press("d")
    time.sleep(presskey)
    keyboard.release("d") 
