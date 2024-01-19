import requests
from bs4 import BeautifulSoup
import tkinter as tk
from tkinter import messagebox

def scrape_website(url):
    try:
        response = requests.get(url)
        response.raise_for_status()
        soup = BeautifulSoup(response.text, 'html.parser')
        # Do something with the soup object (extract data, etc.)
        # For now, let's just print the title
        title = soup.title.string
        messagebox.showinfo('Scraped Content', f'Title: {title}')
    except requests.exceptions.RequestException as e:
        messagebox.showerror('Error', f'Failed to scrape website: {e}')

def on_submit():
    url = entry.get()
    scrape_website(url)

# GUI setup
root = tk.Tk()
root.title('Web Scraper')

label = tk.Label(root, text='Enter website URL:')
label.pack(pady=10)

entry = tk.Entry(root, width=40)
entry.pack(pady=10)

submit_button = tk.Button(root, text='Scrape', command=on_submit)
submit_button.pack(pady=10)

root.mainloop()
