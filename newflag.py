import requests, datetime, argparse
from bs4 import BeautifulSoup

def banner():
    print("NewRedflag --- by 0xlildoudou\n")

def download_all(link):
    new_redflagdomain = requests.get('https://dl.red.flag.domains/daily/' + link.get('href')).text
    if output:
        filename = output
    else:
        filename = "redflag_all.txt"
    file = open(filename, "a")
    file.write(new_redflagdomain)
    file.close()
    return filename

def download_date(date, link):
    new_redflagdomain = requests.get('https://dl.red.flag.domains/daily/' + link.get('href')).text
    if output:
        filename = output
    else:
        filename = "redflag_" + str(date) + ".txt"
    file = open(filename, "a")
    file.write(new_redflagdomain)
    file.close()
    return filename

def main(latest, day, all):
    if not day or latest == True:
        date = datetime.datetime.now() - datetime.timedelta(days=1)
        yesterday_date = date.strftime("%Y-%m-%d")
    else:
        date = datetime.datetime.now() - datetime.timedelta(days=int(day))
        specific_date = date.strftime("%Y-%m-%d")
    main_url = requests.get('https://dl.red.flag.domains/daily/').text
    main_page = BeautifulSoup(main_url, 'html.parser')
    for link in main_page.find_all('a'):
        date_url = link.get('href')
        if date_url[:4].isdigit():
            try:
                if (yesterday_date in date_url) and (latest == True):
                    if verbose:
                        print("[Download] : " + str(date_url.removesuffix('.txt')))
                    output_file =  download_date(yesterday_date, link)
            except:
                pass

            try:
                if (specific_date in date_url) and day:
                    if verbose:
                        print("[Download] : " + str(date_url.removesuffix('.txt')))
                    output_file = download_date(specific_date, link)
            except:
                pass

            if all == True:
                if verbose:
                    print("[Download] : " + str(date_url.removesuffix('.txt')))
                output_file = download_all(link)
    
    print("---> " + output_file)
            

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--latest","-l", help="latest list", action="store_true")
    parser.add_argument("--day","-d", help="Day of list")
    parser.add_argument("--all","-A", help="All list", action="store_true")
    parser.add_argument("--output","-o", help="Output file")
    parser.add_argument("--verbose", "-v", help="Output file", action="store_true")
    args = parser.parse_args()

    latest = args.latest
    day = args.day
    all = args.all

    global output
    global verbose
    output = args.output
    verbose = args.verbose

    banner()
    main(latest, day, all)