import pandas as pd
import numpy as np
from IPython.display import display
from tabulate import tabulate

df = pd.DataFrame(
    {"A": [1, 2, 3, 4, 5], "B": [10, 20, 30, 40, 50]}, index=["a", "b", "c", "d", "e"]
)

print(df.shape)

# show all data
print("show All")
display(df)

print("-------\n")
print("show A")
# show A Col's data
display(df["A"])

print("-------\n")
print("show rearrange")
rearrange = pd.DataFrame(df, columns=["B", "A"])
display(rearrange)

# print('-------\n')
# print('select a rows')
# display(df.loc('a'))
dict = {
    "Name": ["Martha", "Tim", "Rob", "Georgia"],
    "Maths": [87, 91, 97, 95],
    "Science": [83, 99, 84, 76],
}
df = pd.DataFrame(dict)

# displaying the DataFrame
df.style


def color_negative_red(val):
    """
    Takes a scalar and returns a string with
    the css property `'color: red'` for negative
    strings, black otherwise.
    """
    color = "blue" if val > 90 else "black"
    return "color: % s" % color


# creating a DataFrame
dict = {"Maths": [87, 91, 97, 95], "Science": [83, 99, 84, 76]}
df = pd.DataFrame(dict)

# displaying the DataFrame
df.style.applymap(color_negative_red)


print(tabulate(df, headers="keys", tablefmt="psql"))

print("*************")
ts = pd.Series(np.random.randn(1000), index=pd.date_range("1/1/2000", periods=1000))
ts = ts.cumsum()
ts.plot()

df = pd.DataFrame(np.random.randn(1000, 4), index=ts.index, columns=list("ABCD"))
df = df.cumsum()
df.plot()
