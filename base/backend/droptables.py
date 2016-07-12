import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'foodies', host = 'localhost', user = 'anil', password = '205474')
cursor = postgres.cursor()


cursor.execute("""
   select exists(select 0 from pg_class where relname = 'vendors')
""")

presence = cursor.fetchone()[0]

if presence:
    print('vendors table exists, deleting here')
    cursor.execute("""
        drop table vendors
    """)


cursor.execute("""
   select exists(select 0 from pg_class where relname = 'itemmenu')
""")

presence = cursor.fetchone()[0]

if presence:
    print('ItemsMenu table exists, deleting here')
    cursor.execute("""
        drop table itemmenu
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'cutomers')
""")

presence = cursor.fetchone()[0]

if presence:
    print('cutomers table exists, deleting here')
    cursor.execute("""
        drop table cutomers
    """)


cursor.execute("""
   select exists(select 0 from pg_class where relname = 'foodiesrecord')
""")

presence = cursor.fetchone()[0]

if presence:
    print('foodiesrecord table exists, deleting here')
    cursor.execute("""
        drop table foodiesrecord
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'accounts_record')
""")

presence = cursor.fetchone()[0]

if presence:
    print('accounts_record table exists, deleting here')
    cursor.execute("""
        drop table accounts_record
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'ordersrecord')
""")

presence = cursor.fetchone()[0]
if presence:
    print(' table exists, deleting here')
    cursor.execute("""
        drop table ordersrecord
    """)

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


