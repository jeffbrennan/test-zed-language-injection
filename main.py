import sqlite3

import pandas as pd
from psycopg2 import sql
from pyspark.sql import SparkSession
from sqlalchemy import create_engine, text

engine = create_engine("sqlite+pysqlite:///:memory:")
con = sqlite3.connect(":memory:")
spark = SparkSession.builder.getOrCreate()

# .sql (polars, pyspark, duckdb)
df = spark.sql("SELECT col1, col2 FROM tbl")
spark.read.text("catalog.schema.table")  # no false positive on .text()

# .SQL (psycopg2)
query = sql.SQL("SELECT col1, col2 FROM tbl")

# .read_sql/.read_sql_query (pandas)
pd.read_sql_query("SELECT col1, col2 FROM tbl", con)
pd.read_sql("SELECT col1, col2 FROM tbl", con)

# execute (sqlite)
with sqlite3.connect(":memory:") as con:
    cursor = con.cursor()
    cursor.execute("SELECT * FROM tbl")

# text (sqlalchemy)
with engine.connect() as connection:
    result = connection.execute(
        text("SELECT col1, col2 FROM tbl")
    )  # text() works here

# string vars
# sql
cmd = "SELECT col1, col2 FROM tbl"

# sql
cmd = """
    SELECT
        col1,
        col2
    FROM tbl
"""

# do not inject sql
cmd = "SELECT col1, col2 FROM tbl"

# abc sql def
cmd = "SELECT col1, col2 FROM tbl"

# sql comment that should not be injected
cmd = "SELECT col1, col2 FROM tbl"
