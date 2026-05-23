# update the server using the following process:
# - Backup userdata.db (for user data)
# - Backup serverdata.db (for events state)
# - Update Elichika source code
# - Restore userdata.db and serverdata.db
# - Rebuild serverdata.db to new state
# running this command also potentially remove outdated backup

mv -f data/userdata.db backup/userdata.db.temp
mv -f data/serverdata.db backup/serverdata.db.temp
mv -f data/config.json backup/config.json.temp
echo "Backed up databases, updating"
git reset --hard HEAD
git checkout master && git pull
cp backup/userdata.db.temp data/userdata.db
cp backup/serverdata.db.temp data/serverdata.db
cp backup/config.json.temp data/config.json
./elichika rebuild_assets
echo "Updated successfully!"


if [ $? -eq 0 ]; then
    chmod +rwx ./bin/shortcut.sh && \
    ./bin/shortcut.sh
else
    echo "Error updating!"
fi