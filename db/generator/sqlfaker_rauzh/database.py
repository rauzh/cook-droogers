from generator.sqlfaker_rauzh.table import Table
from generator.sqlfaker_rauzh.functions import list_to_string, check_type
from generator.sqlfaker_rauzh.sql_dictionary import sql_dictionary


error_01 = "Parameter {} was {}, but can only be: {}"


class Database:
    """This is the main class of this package. It is used to instantiate database objects.

    A database object holds multiple tables which again can hold multiple
    columns. Every data structure, that is done with generator.sqlfaker_rauzh starts 
    with a database.

    Currently, generator.sqlfaker_rauzh only supports mysql/mariadb syntax. The db_type parameter
    is therefore useless at the moment and will be set to mysql by default.
    
    :param db_type: The type of database to export SQL for (generator.sqlfaker_rauzh currently only supports mysql/mariadb)
    :param db_name: The name of the database
    :type db_name: String
    :type db_type: String
    """

    def __init__(self, db_name, dbs_type="mysql", lang="de_DE"):
        
        # check if dbs_type is allowed
        allowed_dbs_types = ["mysql", "mariadb", "sqlite"]
        if dbs_type not in allowed_dbs_types:
            raise ValueError(
                error_01.format(
                    "dbs_type",
                    dbs_type,
                    list_to_string(allowed_dbs_types)
                )
            )
        
        # Store parameters in object
        self._db_name = db_name
        self._type = dbs_type

        # Add room for all tables of this table
        self.tables = {}
        self.lang = lang

    def add_table(self, table_name, n_rows):
        """This method can be used to add a table to the database.

        The table object will be stored in the tables dictionary of the
        database. The table name will be used as key.

        :param table_name: Name of the new table
        :type table_name: String
        :param n_rows: Number of rows that the table should have in DML
        :type n_rows: Integer
        :returns: None
        :raises ValueError: If n_rows is not integer
        :raises ValueError: If table_name is not string
        :raises ValueError: IF n_rows is 0 or less
        """

        # raise error if types do not match
        check_type(n_rows, int)
        check_type(table_name, str)
        if n_rows < 1:
            raise ValueError("n_rows must be at least 1 but was {}".format(
                str(n_rows)
            ))

        self.tables[table_name] = Table(
            table_name=table_name,
            db_object=self,
            n_rows=n_rows
        )

    def generate_data(self, recursive=False):
        """This method runs all generator methods of all tables
        
        To do so, the method will iterate all stored table objects and
        will run the generate_data method of each table.

        :param recursive: Wether data generation is done for rekursive data
        :type recursive: Boolean
        :default recursive: False
        :returns: None
        """
        for key in self.tables.keys():
            self.tables[key].generate_data(recursive=recursive, lang=self.lang)


    # ##############################################
    # Getter and Setters
    # ##############################################
    
    def return_db_name(self):
        """This method returns the name of a database object.
        
        :returns: Database name
        """
        return self._db_name

    def return_db_tables(self):
        """This method returns the names of all stored table objects.
        
        :returns: List of table names
        """

        return self.tables.keys()

    def return_ddl(self):
        """This method generates the database's DDL and returns it as string.
        
        :returns: DDL script as string
        """
        
        # TODO DDL output must be adopted (maybe a dictionary going by the name
        # of the dbs as key...)
        
        ddl_output = "DROP DATABASE IF EXISTS {};\n".format(self._db_name)
        ddl_output += "CREATE DATABASE {};\n".format(self._db_name)
        ddl_output += "USE {};\n\n".format(self._db_name)

        for key in self.tables:
            table_object = self.tables[key]
            
            ddl_output += table_object.return_ddl()
        
        return ddl_output

    def return_dml(self):
        """This method generates the database's DML and returns it as string.
        
        :returns: DDL script as string
        """
        
        # TODO This also needs to be adopted for dbs support

        dml_output = "USE {};\n\n".format(self._db_name)

        for key in self.tables:
            table_object = self.tables[key]
            
            dml_output += table_object.return_dml()
        
        return dml_output    

    def export_ddl(self, file_name):
        """This method exports the database's DDL script to disk.
        
        :param file_name: The file name (e.g. "C:/my_ddl.sql")
        :type file_name: String
        :returns: None
        """

        with open(file_name, "w", encoding="utf-8") as out_file:
            out_file.write(self.return_ddl())

    def export_dml(self, file_name):
        """This method exports the database's DML script to disk.
        
        :param file_name: The file name (e.g. "C:/my_ddl.sql")
        :type file_name: String
        :returns: None
        """

        with open(file_name, "w", encoding="utf-8") as out_file:
            out_file.write(self.return_dml())
    
    def export_sql(self, file_name):
        """This method exports the database's complete SQL script to disk.
        
        :param file_name: The file name (e.g. "C:/my_ddl.sql")
        :type file_name: String
        :returns: None
        """

        with open(file_name, "w", encoding="utf-8") as out_file:
            out_file.write(self.return_ddl())
            
        with open(file_name, "a", encoding="utf-8") as out_file:
            out_file.write(self.return_dml())
