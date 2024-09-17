from generator.sqlfaker_rauzh.database import Database

def add_users(db, n = 100):

    db.add_table(table_name="users", n_rows=n)

    db.tables["users"].add_primary_key(column_name="user_id")
    db.tables["users"].add_column(column_name="name", data_type="text", data_target="name")
    db.tables["users"].add_column(column_name="email", data_type="text", data_target="email")
    db.tables["users"].add_column(column_name="password", data_type="text", data_target="password")
    db.tables["users"].add_column(column_name="type", data_type="int", column_value=0)

    db.tables["users"].generate_data(recursive=False, lang="en_US")

def add_managers(db, n = 10):
    
    db.add_table(table_name="managers", n_rows=n)

    db.tables["managers"].add_primary_key(column_name="manager_id") 
    db.tables["managers"].add_foreign_key(column_name="user_id", target_table="users", target_column="user_id")

    db.tables["managers"].generate_data(recursive=False, lang="en_US")

def add_artists(db, n = 50):
    
    db.add_table(table_name="artists", n_rows=n)

    db.tables["artists"].add_primary_key(column_name="artist_id") 

    db.tables["artists"].add_foreign_key(column_name="manager_id", target_table="managers", target_column="manager_id") 
    db.tables["artists"].add_foreign_key(column_name="user_id", target_table="users", target_column="user_id")

    db.tables["artists"].add_column(column_name="nickname", data_type="text", data_target="user_name")
    db.tables["artists"].add_column(column_name="contract_due", data_type="timestamp", data_target="future_date")
    db.tables["artists"].add_column(column_name="activity", data_type="boolean", data_target="boolean")

    db.tables["artists"].generate_data(recursive=False, lang="en_US")

def add_releases(db, n = 50):

    db.add_table(table_name="releases", n_rows=n)

    db.tables["releases"].add_primary_key(column_name="release_id") 

    db.tables["releases"].add_foreign_key(column_name="artist_id", target_table="artists", target_column="artist_id") 

    db.tables["releases"].add_column(column_name="title", data_type="text", data_target="company")
    db.tables["releases"].add_column(column_name="status", data_type="text", column_value="Unpublished")
    db.tables["releases"].add_column(column_name="creation_date", data_type="timestamp", data_target="past_date")

    db.tables["releases"].generate_data(recursive=False, lang="en_US")

def add_publications(db, n = 30):

    db.add_table(table_name="publications", n_rows=n)

    db.tables["publications"].add_primary_key(column_name="publication_id") 

    db.tables["publications"].add_foreign_key(column_name="manager_id", target_table="managers", target_column="manager_id") 
    db.tables["publications"].add_foreign_key(column_name="release_id", target_table="releases", target_column="release_id")

    db.tables["publications"].add_column(column_name="creation_date", data_type="timestamp", data_target="past_date")

    db.tables["publications"].generate_data(recursive=False, lang="en_US")

def add_tracks(db, n = 300):

    db.add_table(table_name="tracks", n_rows=n)

    db.tables["tracks"].add_primary_key(column_name="track_id") 

    db.tables["tracks"].add_foreign_key(column_name="release_id", target_table="releases", target_column="release_id")

    db.tables["tracks"].add_column(column_name="title", data_type="text", data_target="company")
    db.tables["tracks"].add_column(column_name="genre", data_type="text", data_target="city")
    db.tables["tracks"].add_column(column_name="type", data_type="text", data_target="country")
    db.tables["tracks"].add_column(column_name="duration", data_type="int", column_value=120)

    db.tables["tracks"].generate_data(recursive=False, lang="en_US")

def add_track_artist(db, n = 300):

    db.add_table(table_name="track_artist", n_rows=n)

    db.tables["track_artist"].add_primary_key(column_name="track_artist_id") 

    db.tables["track_artist"].add_foreign_key(column_name="track_id", target_table="tracks", target_column="track_id")
    db.tables["track_artist"].add_foreign_key(column_name="artist_id", target_table="artists", target_column="artist_id")

    db.tables["track_artist"].generate_data(recursive=False, lang="en_US")

db = Database(db_name="public")
add_users(db, 10000*5)
add_managers(db, 1000*5)
add_artists(db, 5000*5)
add_releases(db, 5000*5)
add_publications(db, 2000*5)
add_tracks(db, 30000*5)
add_track_artist(db, 30000*5)

# user-manager should have type 1
for i in range(db.tables["managers"]._n_rows):
    user_id = db.tables['managers'].columns['user_id'].data[i]
    db.tables['users'].columns['type'].data[user_id - 1] = 1

# user-artist should have type 2
for i in range(db.tables["artists"]._n_rows):
    user_id = db.tables['artists'].columns['user_id'].data[i]
    db.tables['users'].columns['type'].data[user_id - 1] = 2 if db.tables['users'].columns['type'].data[user_id - 1] == 0 else 1

# publications are published releases
for i in range(db.tables["publications"]._n_rows):
    release_id = db.tables['publications'].columns['release_id'].data[i]
    db.tables['releases'].columns['status'].data[release_id - 1] = 'Published'

# track author is release's owner
for i in range(db.tables["track_artist"]._n_rows):
    track_id = db.tables['track_artist'].columns['track_id'].data[i]
    release_id = db.tables['tracks'].columns['release_id'].data[track_id - 1]
    artist_id = db.tables['releases'].columns['artist_id'].data[release_id - 1]
    db.tables['track_artist'].columns['artist_id'].data[i] = artist_id

# publication'a manager is authors's manager
for i in range(db.tables["publications"]._n_rows):
    release_id = db.tables['publications'].columns['release_id'].data[i]
    artist_id = db.tables['releases'].columns['artist_id'].data[release_id - 1]
    manager_id = db.tables['artists'].columns['manager_id'].data[artist_id - 1]
    db.tables['publications'].columns['manager_id'].data[i] = manager_id

# db.generate_data()

print(db.tables["users"].return_dml())
print(db.tables["managers"].return_dml())
print(db.tables["artists"].return_dml())
print(db.tables["releases"].return_dml())
print(db.tables["publications"].return_dml())
print(db.tables["tracks"].return_dml())
print(db.tables["track_artist"].return_dml())
