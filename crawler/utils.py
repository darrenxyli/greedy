from colorama import Fore

# @Method: red()
# @Description: return a string with red color


def red(s):
    return Fore.RED + s + Fore.RESET


# @Method: green()
# @Description: return a string with green color
def green(s):
    return Fore.GREEN + s + Fore.RESET
