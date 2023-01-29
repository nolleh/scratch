import pandas
from IPython.display import display

df = pandas.DataFrame(
    {"A": [1, 2, 3, 4, 5], "B": [10, 20, 30, 40, 50]}, index=["a", "b", "c", "d", "e"]
)

print(df.shape)

# show all data
print('show All')
display(df)

print('-------\n')
print('show A')
# show A Col's data
display(df['A'])

print('-------\n')
print('show rearrange')
rearrange = pandas.DataFrame(df, columns = ['B','A'])
display(rearrange)

# print('-------\n')
# print('select a rows')
# display(df.loc('a'))

