import psycopg2
import random
import requests
import json

postgres = psycopg2.connect(database = 'foodies', host = 'localhost', user = 'anil', password = '205474')
cursor = postgres.cursor()


cursor.execute("""
   select exists(select 0 from pg_class where relname = 'accounts_record')
""")

presence = cursor.fetchone()[0]

if presence:
    print('accounts_record table exists')
    cursor.execute("""
        drop table accounts_record
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'ordersrecord')
""")

presence = cursor.fetchone()[0]
if presence:
    print('ordersrecord table exists')
    cursor.execute("""
        drop table ordersrecord
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'foodiesrecord')
""")

presence = cursor.fetchone()[0]

if presence:
    print('foodiesrecord table exists')
    cursor.execute("""
        drop table foodiesrecord
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'itemmenu')
""")

presence = cursor.fetchone()[0]

if presence:
    print('ItemsMenu table exist')
    cursor.execute("""
        drop table itemmenu
    """)

cursor.execute("""
   select exists(select 0 from pg_class where relname = 'vendors')
""")

presence = cursor.fetchone()[0]

if presence:
    print('vendors table exist')
    cursor.execute("""
        drop table vendors
    """)



cursor.execute("""
   select exists(select 0 from pg_class where relname = 'customers')
""")

presence = cursor.fetchone()[0]

if presence:
    print('customers table exist')
    cursor.execute("""
        drop table customers
    """)
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)



#                       CREATE TABLE VENDORS                                         #              
cursor.execute("""
    create table vendors(
        vendorid bigserial PRIMARY KEY,
        owner text NOT NULL,
        vendorname text NOT NULL,
        email text NOT NULL,
        mobile varchar(12)[] NOT NULL,
        address text NOT NULL,
        imageaddress text,
        description text ,
        offer text,
        password text NOT NULL
    )
""")

#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                       CREATE TABLE ITEMSMENU                                      #              
cursor.execute("""
    create table itemmenu(
        vendor_id int NOT NULL references vendors(vendorid),
        item_no bigint NOT NULL ,        
        item_name text NOT NULL,
        item_type varchar(10) NOT NULL CHECK(item_type in ('starter','main','desert')) ,
        item_nature varchar(1) NOT NULL CHECK(item_nature in ('v','n')),
        price text NOT NULL,
        item_description text NOT NULL,
        imageaddress text,
        discount double precision,
        PRIMARY KEY(vendor_id, item_no)
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)



#                       CREATE TABLE CUSTOMERS                                            #
cursor.execute("""
    create table customers(
            customer_id bigserial PRIMARY KEY NOT NULL,
            customer_name text NOT NULL,
            emailid text NOT NULL,
            mobile varchar(12)[] NOT NULL,
            address text,
            password text NOT NULL
    )
""")
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


#                               CHECK FOR TABLE FOODIES_RECORD                      #
# cursor.execute("""
#    select exists(select 0 from pg_class where relname = 'foodiesrecord')
# """)

# presence = cursor.fetchone()[0]
# print(presence)

#  #                       DELETE FOODIESRECORD IF EXISTS                              #
# if presence:
#     print('foodiesrecord table exists, deleting here')
#     cursor.execute("""
#         drop table foodiesrecord
#     """)

# #                       CREATE TABLE FOODIESRECORD                                     #
# cursor.execute("""
#     create table foodiesrecord(
#             transactionid bigserial PRIMARY KEY NOT NULL,
#             ordered_on timestamp NOT NULL,
#             customer_id int NOT NULL,
#             delivered_on timestamp,
#             vendorid int NOT NULL,
#             status varchar(20) NOT NULL,
#             bill money NOT NULL
#     )
# """)
# try:
#     postgres.commit()
# except Exception as e:
#     postgres.rollback()
#     print(e)



#                       CREATE TABLE ACCOUNTS_RECORD                                    #
cursor.execute("""
    create table accounts_record(
        v_id bigserial PRIMARY KEY NOT NULL references vendors(vendorid),
        account smallint NOT NULL
    )
""")
try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)




#                       CREATE TABLE ORDERsRECORD                                   #              
cursor.execute("""
    create table ordersrecord(
        order_id int PRIMARY KEY,
        vendor_id int NOT NULL references vendors(vendorid),
        customer_id int NOT NULL references customers(customer_id),        
        ordered_placed_on timestamp NOT NULL,
        ordered_deliverd_on timestamp ,
        order_status character(1) CHECK(order_status in('y','n')),
        description text NOT NULL,
        amount text NOT NULL
    )
""")
#                       COMMIT AND ROLLBACK IF EXCEPTION                             #

try:
    postgres.commit()
except Exception as e:
    postgres.rollback()
    print(e)


