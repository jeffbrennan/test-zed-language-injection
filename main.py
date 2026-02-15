from pyspark.sql import SparkSession

spark = SparkSession.builder.getOrCreate()

# sql
cmd = "SELECT col1, col2 FROM tbl"


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


my_regex = r"[0-9]"
spark.sql(  # sql
    f"""
        SELECT
        col1,
        CASE WHEN col2 RLIKE '{my_regex}' THEN 1 ELSE 2 END as col4,
        ROW_NUMBER() OVER(PARTITION BY col1, col2 ORDER BY col3) as rn
        FROM base
    )
    """
)
