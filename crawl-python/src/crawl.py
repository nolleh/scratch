import requests as req
from bs4 import BeautifulSoup
from pandas import DataFrame
from IPython.display import display
from tabulate import tabulate

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

rank = df.groupby(['lang'])['lang'].size().sort_values(ascending=False).reset_index(name='counts')

print(tabulate(rank, headers = 'keys', tablefmt='psql'))

f = open('crawl.html', 'w', newline='')
cell_hover = {  # for row hover use <tr> instead of <td>
    'selector': 'td:hover',
    'props': [('background-color', '#ffffb3')]
}
index_names = {
    'selector': '.index_name',
    'props': 'font-style: italic; color: darkgrey; font-weight:normal;'
}
headers = {
    'selector': 'th:not(.index_name)',
    'props': 'background-color: #000066; color: white;'
}

genhtml = rank.style.set_table_styles([cell_hover, index_names, headers]).to_html()
wr = f.write(genhtml) 
f.close()

