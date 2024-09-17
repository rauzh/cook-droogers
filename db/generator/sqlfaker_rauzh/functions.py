from faker import Faker
import datetime

error_02 = "List contains to elements."


def get_fake_data(data_target="name", n_rows=100, lang="de_DE"):
    """This method generates a certain data item based on the data_target selected
    
    :param data_target: Data item type that should be created
    :type data_target: String
    :param n_rows: Number of rows to generate
    :type n_rows: Integer
    :returns: List of data
    """
    data_faker = Faker(lang)
    generator_function = getattr(data_faker, data_target)
    return_list = []
    
    for _ in range(n_rows):
        if data_target in ["date", "past_datetime", "time", "past_date", "date_time"]:
            return_list.append(datetime.datetime.strftime(generator_function(), "%Y-%m-%d %H:%M:%S"))
        else:
            return_list.append(generator_function())
    return return_list


def check_type(value, planned_type=str):
    """This method checks if the handed value's type is equal to planned.
    
    :param value: The value to check
    :type value: Python object
    :param plannd_type: The intended python object type
    :type planned_type: Python type
    :return: None
    :raises ValueError: If types are not equal
    """

    #error_message = "Value of {} was expected to contain {} but was {}."
    #if type(value) != str:
    #        raise ValueError(error_message.format(
    #            namestr(value),
    #            "str",
    #            str(type(value))
    #        ))
    pass


def check_value_is_not_less_than(value, compare_value):
    """This method checks if a value is larger than a certain value.
    
    :param value: Object to check
    :param compare_value: Threshold to compare with
    """
    if value < compare_value:
        raise ValueError("n_rows must be at least {}, but was {}.".format(
            str(compare_value),
            str(value)
        ))

def check_value_is_not_more_than(value, compare_value):
    """This method checks if a value is smaller than a certain value.
    
    :param value: Object to check
    :param compare_value: Threshold to compare with
    """
    if value > compare_value:
        raise ValueError("n_rows must be at least {}, but was {}.".format(
            str(compare_value),
            str(value)
        ))



def namestr(object):
    """This method returns the string name of a python object.
    
    :param object: The object that should be used
    :type object: Python object
    :returns: Object name as string
    """
    for n,v in globals().items():
        if v == object:
            return n
    return None


def list_to_string(value_list):
    """Method concatenates a list of values to comma seperated string
    
    :param list: Python list of values
    :type list: List of strings
    :returns: Comma separated string
    :raises ValueError: If value_list is empty
    :raises ValueError: If value_list entries are not type str
    """

    # check if list contains elements
    if len(value_list) == 0:
        raise ValueError(error_02)

    # check if all values are type str
    for element in value_list:
        check_type(element, str)

    # return result
    return ", ".join(value_list)
