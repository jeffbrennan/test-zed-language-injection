import sqlite3

import pandas as pd
from pyspark.sql import SparkSession
from sqlalchemy import create_engine, text

engine = create_engine("sqlite+pysqlite:///:memory:")
con = sqlite3.connect(":memory:")
spark = SparkSession.builder.getOrCreate()

df = spark.sql("SELECT col1, col2 FROM tbl")
df = spark.sql(  # sql
    "SELECT col1, col2 FROM tbl"
)
df = spark.sql(  # sql
    """
    WITH cte AS (
        SELECT col1, col2, COUNT(*) AS n
        FROM tbl
        GROUP BY ALL
    )
    SELECT * FROM cte
    """
)

# .read_sql/.read_sql_query (pandas)
pd.read_sql_query(  # sql
    "SELECT col1, col2 FROM tbl", con
)
pd.read_sql("SELECT col1, col2 FROM tbl", con)

# execute (sqlite)
with sqlite3.connect(":memory:") as con:
    cursor = con.cursor()
    cursor.execute(
        # sql
        "SELECT * FROM tbl"
    )
    cursor.execute("SELECT * FROM tbl")

# text (sqlalchemy)
with engine.connect() as connection:
    result = connection.execute(
        text(
            # sql
            "SELECT col1, col2 FROM tbl"
        )
    )

# string vars
# sql
cmd = "SELECT col1, col2 FROM tbl"

# sql
cmd = """
    SELECT col1, col2
    FROM tbl
"""

# do not inject sql
cmd = "SELECT col1, col2 FROM tbl"

# abc sql def
cmd = "SELECT col1, col2 FROM tbl"

# sql comment that should not be injected
cmd = "SELECT col1, col2 FROM tbl"
