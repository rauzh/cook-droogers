from generator.sqlfaker_rauzh.functions import get_fake_data, check_type


class Column:
    """The column class represents column objects of a table.
    
    This class is used to setup a new column, define its meta data and generate
    column data afterwards. Data are stored within the column object and are later
    used when exporting the DML script for the table.

    :param column_name: Defines the colum's name
    :param n_rows: Defines how many rows to generate for the column
    :param data_type: Defines the columns data type
    :param ai: Defines the auto increment attribute of the column
    :param not_null: Defines the not null attribute of the column
    :param data_target: The type of fake data that faker should create
    :param value: The only value that should be in this column
    
    :type column_name: String
    :type n_rows: Integer
    :type ai: Boolean
    :type not_null: Boolean
    :type data_type: String
    :type data_target: String
    :type value: Same as data_type

    :raises ValueError: If column_name not string
    :raises ValueError: If data_target not string
    :raises ValueError: If data_type not string
    :raises ValueError: If ai not boolean
    :raises ValueError: If not_null not boolean
    """
    
    def __init__(
        self,
        column_name,
        n_rows,
        table_objet,
        data_target="name",
        data_type="int",
        ai=False,
        not_null=False,
        value=None
    ):
        
        # Check type of input data
        check_type(column_name, str)
        check_type(data_target, str)
        check_type(data_type, str)
        check_type(ai, bool)
        check_type(not_null, bool)

        # store all parameters
        self._column_name = column_name
        self._data_type = data_type
        self._ai = ai
        self._not_null = not_null
        self._n_rows = n_rows
        self._data_target = data_target

        self.value = value

        # store own table object
        self._table_object = table_objet

        # store data
        self.data = []

    def generate_data(self, recursive, lang):
        """This method generates data for a column object.
        
        :param recursive: Wether data generation is done for rekursive data
        :type recursive: Boolean
        :returns: None
        """

        if self._ai == True:
            
            # generate incrementing values from 1 to n
            self.data = list(range(1, self._n_rows+1))

        else:

            if self.value is not None:
                self.data = [self.value for _ in range(self._n_rows)]
                return

            # generate data using faker
            self.data = get_fake_data(
                data_target=self._data_target,
                n_rows=self._n_rows,
                lang=lang
            )

    def return_ddl(self):
        """This method returns the DDL line of the respective column.
        
        :returns: DDL line as String
        """

        # TODO Adopt this for dbs support
        
        name = self._column_name
        not_null = self._not_null
        ai = self._ai
        data_type = str.upper(self._data_type)

        ddl_output = "\t{} {}{}{},\n".format(
            name,
            data_type,
            " AUTO_INCREMENT" if ai else "",
            " NOT NULL" if not_null else ""
        )

        return ddl_output
