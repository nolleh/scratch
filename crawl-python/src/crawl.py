import requests as req
from bs4 import BeautifulSoup
from pandas import DataFrame
from IPython.display import display

url = 'https://github.com/nolleh?tap=repositories'
res = req.get(url, params = {'query' : ''})

soup = BeautifulSoup(res.text, 'html.parser')
langs = soup.find_all(itemprop='programmingLanguage')

# print(langs)
df = DataFrame(langs, columns=['lang'])
display(df)

print('********************************')
print('** group by and column count by columns **')

display(df.groupby(['lang'])['lang'].count())

print('********************************')
print('** group by and rename index with counts **')
display(df.groupby(['lang'])['lang'].size().reset_index(name='counts'))

print('********************************')
print('** group by and sort with count **')

display(df.groupby(['lang'])['lang'].size().sort_values(ascending=False))


