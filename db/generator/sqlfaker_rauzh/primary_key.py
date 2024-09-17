from generator.sqlfaker_rauzh.column import Column


class PrimaryKey(Column):
    """This column class represents primary key objects of a table.
    
    This class is used to setup a new column that is used as an identifier
    in a table. It is a sub class of Column and inherits most of its
    attributes.

    :param column_name: Defines the colum's name
    :param n_rows: Defines how many rows to generate for the column
    :param table_object: Holds the parent table which the key belongs to
    :type column_name: String
    :type table_object: generator.sqlfaker_rauzh Table object
    :type n_rows: Integer
    """

    def __init__(self, column_name, n_rows, table_objet):
        
        # Instantiate the master class but fix some parameters
        super().__init__(
            column_name = column_name,
            n_rows=n_rows,
            ai=True,
            not_null=True,
            table_objet=table_objet,
            data_target=None,
            data_type="int"
        )

    def return_ddl(self):
        """This method returns the DDL line of the respective key column.
        
        :returns: DDL line as String
        """

        ddl_output = "\t{} INT PRIMARY KEY AUTO_INCREMENT NOT NULL,\n".format(
            self._column_name
        )

        return ddl_output

