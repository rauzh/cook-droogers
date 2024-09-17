from generator.sqlfaker_rauzh.column import Column
from generator.sqlfaker_rauzh.foreign_key import ForeignKey
from generator.sqlfaker_rauzh.primary_key import PrimaryKey
from generator.sqlfaker_rauzh.functions import check_type, check_value_is_not_less_than
from numpy import array


class Table:
    """The table class lets you add column families to a database.
    
    A table is a collection of columns of which some are primary keys,
    others are foreign keys and some might be regular columns. For the
    sake of fake data creation, you have to select a number of rows
    that should be created for this table.

    :param table_name: Name of the table to be created
    :param db_object: The database, the table belongs to
    :param n_rows: The number of rows that should be created for DML
    :type table_name: String
    :type db_object: Python generator.sqlfaker_rauzh Database object
    :type n_rows: Integer
    :raises ValueError: If n_rows is smaller than 1
    :raises ValueError: If n_rows is not integer
    """
    
    def __init__(self, table_name, db_object, n_rows=100):

        # check inputs
        check_type(n_rows, int)
        check_value_is_not_less_than(n_rows, 1)
        
        # Store parameters in object
        self._table_name = table_name
        self._n_rows = n_rows

        # store own database object
        self._db_object = db_object

        # Add room for all columns of this table
        self.columns = {}


    def add_column(self, column_name, data_type="int", not_null=False, data_target="name", column_value=None):
        """This method adds a new column to a table.
        
        :param column_name: The column's name
        :type column_name: String
        :param data_type: The sql data type that should be used in DDL
        :type data_type: String
        :param not_null: Flag indicating if column is NOT NULL
        :type not_null: Boolean
        :param data_target: The type of data that should be created by faker
        :type data_target: String
        :param column_value: The value that should be set in whole column
        :type data_target: same as data_type
        """

        self.columns[column_name] = Column(

            # add column properties
            column_name= column_name,
            data_type= data_type,
            ai= False,
            not_null= not_null,
            data_target=data_target,

            value=column_value,

            # auto add table properties
            n_rows=self._n_rows,
            table_objet=self
        )

    def add_foreign_key(self, column_name, target_table, target_column):
        """This method adds a foreign key column to a table.

        :param column_name: Name of the foreign key to add
        :param target_table: Name of referenced table
        :param target_column: Name of referenced column
        :type column_name: String
        :type target_table: String
        :type target_column: String
        """

        self.columns[column_name] = ForeignKey(
            
            # add foreign key properties
            column_name=column_name,
            target_table=target_table,
            target_column=target_column,

            # auto add table properties
            n_rows=self._n_rows,
            target_db=self._db_object,
            table_objet=self
        )

    def add_primary_key(self, column_name):
        """This method adds a primary key column to a table.
        
        :param column_name: Name of the foreign key to add
        """

        self.columns[column_name] = PrimaryKey(
            
            # add foreign key properties
            column_name=column_name,

            # auto add table properties
            n_rows=self._n_rows,
            table_objet=self
        )

    def generate_data(self, recursive, lang):
        """This method iterates all columns and calls their data generation method.
        
        :param recursive: Wether data generation is done for rekursive data
        :type recursive: Boolean
        :returns: None
        """
        for key in self.columns.keys():
            self.columns[key].generate_data(recursive=recursive, lang=lang)

    def return_ddl(self):
        """This method returns the DDL script of a table.
        
        :returns: DDL statement as String
        """

        # TODO Adopt this for dbs support

        ddl_output = "CREATE TABLE {}.{} (\n".format(
            self._db_object._db_name,
            self._table_name
        )

        for key in self.columns:
            ddl_output += self.columns[key].return_ddl()

        # remove the comma at the end of the last line
        ddl_output = ddl_output[:-2]

        # add closing braket
        ddl_output += "\n);\n\n"

        return ddl_output

    def return_dml(self):
        """This method returns a table's DML script.

        :returns: DML statement as String
        """
        data = []
        attributes = []
        datatype = []

        # get all data into one place
        for key in self.columns:
            datatype.append(self.columns[key]._data_type)
            data.append(self.columns[key].data)
            attributes.append(key)

        # transpose the data
        data = array(data).transpose()

        # get table meta
        rows = self._n_rows
        cols = len(attributes)

        # TODO Adopt this for dbs support

        header = "INSERT INTO {}.{} ({}) VALUES\n".format(
            self._db_object._db_name,
            self._table_name, 
            ", ".join(list(attributes))
        )

        dml_output = ""

        numtypes = ["int", "float", "single", "decimal", "numeric"]

        for row in range(rows):
            line = ""
            for col in range(cols):
                if col > 0:
                    line += ", "
                if datatype[col].split("(")[0] not in numtypes:
                    line += "'" + str(data[row][col]) + "'"
                if datatype[col].split("(")[0] in numtypes:
                    line += str(data[row][col])
            dml_output += header + "\t" + line + ");\n"

        # add semi colon to end of statement
        dml_output = dml_output[:-2] + ";\n\n"

        return dml_output