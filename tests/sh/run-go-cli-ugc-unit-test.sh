#!/bin/bash

#Meta:
#- random seed: 256
#- template file: go-cli-unit-test.j2

#Instructions:
#- Run the Justice SDK Mock Server first before running this script.

MODULE='cmd'
MODULE_PATH='../samples/cli'
TEMP_TOKEN="/tmp/justice-sample-apps/userData"
TEMP_FILE='file.tmp'

OPERATIONS_COUNT=84

FINISHED_COUNT=0
SUCCESS_COUNT=0
FAILED_COUNT=0

export JUSTICE_BASE_URL="http://0.0.0.0:8080"
export APP_CLIENT_ID="admin"
export APP_CLIENT_SECRET="admin"

create_file() {
    touch $1
}

delete_file() {
    [ ! -e $1 ] || rm $1
}

update_status() {
    return_code=$1
    operation_name=$2

    FINISHED_COUNT=$(( $FINISHED_COUNT + 1 ))

    if [ $return_code == 0 ]
    then
        SUCCESS_COUNT=$(( $SUCCESS_COUNT + 1 ))
        echo "ok $FINISHED_COUNT $operation_name"
    else
        FAILED_COUNT=$(( $FAILED_COUNT + 1 ))
        echo "not ok $FINISHED_COUNT - $operation_name"
        echo '  ---'
        echo '  error: |-'
        while read line; do
          echo "    $line" | tr '\n' ' '
        echo $line
        done < $TEMP_FILE
    fi
}

create_file 'tmp.dat'

rm -f $TEMP_TOKEN
echo {"\"access_token"\":"\"foo"\"} >> $TEMP_TOKEN
echo "1..$OPERATIONS_COUNT"

#- 1 singleAdminGetChannel
samples/cli/sample-apps Ugc singleAdminGetChannel \
    --namespace 'FtBxyZcD' \
    --limit 'XBpGlsQu' \
    --offset 'Ju8vMf0I' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminGetChannel'
delete_file $TEMP_FILE

#- 2 adminCreateChannel
samples/cli/sample-apps Ugc adminCreateChannel \
    --body '{"name": "sJkTrd8I"}' \
    --namespace 'DcV2zXnT' \
    >$TEMP_FILE 2>&1
update_status $? 'adminCreateChannel'
delete_file $TEMP_FILE

#- 3 singleAdminUpdateChannel
samples/cli/sample-apps Ugc singleAdminUpdateChannel \
    --body '{"name": "KjXY1bPq"}' \
    --channelId 'amiBxx9C' \
    --namespace 's18EY84e' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminUpdateChannel'
delete_file $TEMP_FILE

#- 4 singleAdminDeleteChannel
samples/cli/sample-apps Ugc singleAdminDeleteChannel \
    --channelId 'kItqRzHU' \
    --namespace '1oh570KQ' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminDeleteChannel'
delete_file $TEMP_FILE

#- 5 adminUploadContentDirect
update_status 0 'adminUploadContentDirect (skipped deprecated)'

#- 6 adminUploadContentS3
samples/cli/sample-apps Ugc adminUploadContentS3 \
    --body '{"contentType": "BVaewc72", "fileExtension": "krSha68n", "name": "3Ynozp1C", "preview": "2KmIQTuB", "subType": "dNEUsxFb", "tags": ["8CJ17M7D"], "type": "JZaMSxEC"}' \
    --channelId 'bZbygyoa' \
    --namespace 'rORoeNHS' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUploadContentS3'
delete_file $TEMP_FILE

#- 7 singleAdminUpdateContentS3
samples/cli/sample-apps Ugc singleAdminUpdateContentS3 \
    --body '{"contentType": "b8Rh3kgs", "fileExtension": "9qqJbnQs", "name": "oBgiVpP8", "preview": "Cm3yvASU", "subType": "oxdxxFqm", "tags": ["AGTJ8IEd"], "type": "agEtp4w2"}' \
    --channelId '9KOu9c19' \
    --contentId 'R6XDqWHk' \
    --namespace 'kP8npLEK' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminUpdateContentS3'
delete_file $TEMP_FILE

#- 8 adminSearchChannelSpecificContent
samples/cli/sample-apps Ugc adminSearchChannelSpecificContent \
    --channelId 'MfjiX7jp' \
    --namespace 'kVZk3IaQ' \
    --creator 'YEmqGodO' \
    --isofficial 'EGt9gPOj' \
    --limit '0c6i0Jkv' \
    --name 'Ias73ucY' \
    --offset 'nFAJ3DK5' \
    --orderby 'T4Eogg0Y' \
    --sortby '39UoYlpv' \
    --subtype '5bVAgtsD' \
    --tags 'hUTDUscb' \
    --type 'QDjbTQuP' \
    --userId 'Mz2PTRlk' \
    >$TEMP_FILE 2>&1
update_status $? 'adminSearchChannelSpecificContent'
delete_file $TEMP_FILE

#- 9 singleAdminUpdateContentDirect
update_status 0 'singleAdminUpdateContentDirect (skipped deprecated)'

#- 10 singleAdminDeleteContent
samples/cli/sample-apps Ugc singleAdminDeleteContent \
    --channelId 'yU89ZPOw' \
    --contentId '6zPFJ42c' \
    --namespace 'wmzBBSMN' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminDeleteContent'
delete_file $TEMP_FILE

#- 11 singleAdminGetContent
samples/cli/sample-apps Ugc singleAdminGetContent \
    --namespace 'coAAOjKN' \
    --limit 'jfcYHm09' \
    --offset '3aYgBU1s' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminGetContent'
delete_file $TEMP_FILE

#- 12 adminSearchContent
samples/cli/sample-apps Ugc adminSearchContent \
    --namespace 'qjyK0XH4' \
    --creator '5PaRSOFQ' \
    --isofficial 'Btu23REZ' \
    --limit '8hRVX7LG' \
    --name 'OvDdYiQS' \
    --offset '9i7mV1C9' \
    --orderby '1pjG9gpx' \
    --sortby 'L6ycTQdv' \
    --subtype 'ln2LAuSQ' \
    --tags 'WEXL6LFE' \
    --type '1YHo9m12' \
    --userId '6ZWc8hHt' \
    >$TEMP_FILE 2>&1
update_status $? 'adminSearchContent'
delete_file $TEMP_FILE

#- 13 adminGetSpecificContent
samples/cli/sample-apps Ugc adminGetSpecificContent \
    --contentId 'WvbNYqgU' \
    --namespace 'qslArFPi' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetSpecificContent'
delete_file $TEMP_FILE

#- 14 adminDownloadContentPreview
samples/cli/sample-apps Ugc adminDownloadContentPreview \
    --contentId 'HUIvaCv8' \
    --namespace 'kU9dBBpd' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDownloadContentPreview'
delete_file $TEMP_FILE

#- 15 adminUpdateScreenshots
samples/cli/sample-apps Ugc adminUpdateScreenshots \
    --body '{"screenshots": [{"description": "sJLhsVyE", "screenshotId": "xrkxoot0"}]}' \
    --contentId 'B7WOferc' \
    --namespace 'ZdpMci37' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateScreenshots'
delete_file $TEMP_FILE

#- 16 adminUploadContentScreenshot
samples/cli/sample-apps Ugc adminUploadContentScreenshot \
    --body '{"screenshots": [{"contentType": "Ds7YSfEx", "description": "aI3uzLte", "fileExtension": "bmp"}]}' \
    --contentId 'bFAlt4hr' \
    --namespace '7HmOYiBA' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUploadContentScreenshot'
delete_file $TEMP_FILE

#- 17 adminDeleteContentScreenshot
samples/cli/sample-apps Ugc adminDeleteContentScreenshot \
    --contentId '5ltAOXml' \
    --namespace 'G6eh1dTd' \
    --screenshotId 'oTFpBIcu' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteContentScreenshot'
delete_file $TEMP_FILE

#- 18 singleAdminGetAllGroups
samples/cli/sample-apps Ugc singleAdminGetAllGroups \
    --namespace 'C1dQY93O' \
    --limit 'JnJ6Te9v' \
    --offset 'D8ldz7Hu' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminGetAllGroups'
delete_file $TEMP_FILE

#- 19 adminCreateGroup
samples/cli/sample-apps Ugc adminCreateGroup \
    --body '{"contents": ["8AD79kdW"], "name": "unvizU0q"}' \
    --namespace '1pHyhhER' \
    >$TEMP_FILE 2>&1
update_status $? 'adminCreateGroup'
delete_file $TEMP_FILE

#- 20 singleAdminGetGroup
samples/cli/sample-apps Ugc singleAdminGetGroup \
    --groupId 'oGgdrysM' \
    --namespace 'izBGSRdP' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminGetGroup'
delete_file $TEMP_FILE

#- 21 singleAdminUpdateGroup
samples/cli/sample-apps Ugc singleAdminUpdateGroup \
    --body '{"contents": ["2l7DNSZ8"], "name": "Aq0XiPLQ"}' \
    --groupId 'XSe07Zdd' \
    --namespace 'OGTMlJjB' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminUpdateGroup'
delete_file $TEMP_FILE

#- 22 singleAdminDeleteGroup
samples/cli/sample-apps Ugc singleAdminDeleteGroup \
    --groupId 'wj9HJHQK' \
    --namespace 'seEdSXRD' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminDeleteGroup'
delete_file $TEMP_FILE

#- 23 singleAdminGetGroupContents
samples/cli/sample-apps Ugc singleAdminGetGroupContents \
    --groupId 'Svguauw1' \
    --namespace 'xT7eMwSl' \
    --limit '9MLH0NnT' \
    --offset 'J2ulNzBv' \
    >$TEMP_FILE 2>&1
update_status $? 'singleAdminGetGroupContents'
delete_file $TEMP_FILE

#- 24 adminGetTag
samples/cli/sample-apps Ugc adminGetTag \
    --namespace 'wJaQa547' \
    --limit 'JllvA8RW' \
    --offset 'SpabUt7x' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetTag'
delete_file $TEMP_FILE

#- 25 adminCreateTag
samples/cli/sample-apps Ugc adminCreateTag \
    --body '{"tag": "k6QxyWhf"}' \
    --namespace 'qoWfJw2o' \
    >$TEMP_FILE 2>&1
update_status $? 'adminCreateTag'
delete_file $TEMP_FILE

#- 26 adminUpdateTag
samples/cli/sample-apps Ugc adminUpdateTag \
    --body '{"tag": "8oWUqvPC"}' \
    --namespace 'Z2HzT7NX' \
    --tagId 'mWDlXsuN' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateTag'
delete_file $TEMP_FILE

#- 27 adminDeleteTag
samples/cli/sample-apps Ugc adminDeleteTag \
    --namespace 'IdQJR5ls' \
    --tagId 'NOlvkfwa' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteTag'
delete_file $TEMP_FILE

#- 28 adminGetType
samples/cli/sample-apps Ugc adminGetType \
    --namespace 'SbnsuLCg' \
    --limit 'ToxuVTek' \
    --offset 'Jgvg6h5H' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetType'
delete_file $TEMP_FILE

#- 29 adminCreateType
samples/cli/sample-apps Ugc adminCreateType \
    --body '{"subtype": ["IpH0Dvip"], "type": "lEk4vj3L"}' \
    --namespace 'Dp4yqDt8' \
    >$TEMP_FILE 2>&1
update_status $? 'adminCreateType'
delete_file $TEMP_FILE

#- 30 adminUpdateType
samples/cli/sample-apps Ugc adminUpdateType \
    --body '{"subtype": ["QUZDpxlH"], "type": "asinGcjr"}' \
    --namespace 'kmRMttgj' \
    --typeId 'DSaIVBmf' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateType'
delete_file $TEMP_FILE

#- 31 adminDeleteType
samples/cli/sample-apps Ugc adminDeleteType \
    --namespace 't3Udg7p9' \
    --typeId 'PGmY2H5k' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteType'
delete_file $TEMP_FILE

#- 32 adminGetChannel
samples/cli/sample-apps Ugc adminGetChannel \
    --namespace 'X4MsisSX' \
    --userId '28nARxWR' \
    --limit 'pv5ou5xt' \
    --offset 'vd28OUfC' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetChannel'
delete_file $TEMP_FILE

#- 33 adminDeleteAllUserChannels
samples/cli/sample-apps Ugc adminDeleteAllUserChannels \
    --namespace 't8UJC5fl' \
    --userId 'Nyj6HsTt' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteAllUserChannels'
delete_file $TEMP_FILE

#- 34 adminUpdateChannel
samples/cli/sample-apps Ugc adminUpdateChannel \
    --body '{"name": "X8P3llna"}' \
    --channelId 'aS9lqyyg' \
    --namespace 'PcfkJIxf' \
    --userId 'QZza8kNV' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateChannel'
delete_file $TEMP_FILE

#- 35 adminDeleteChannel
samples/cli/sample-apps Ugc adminDeleteChannel \
    --channelId 'bDxVMq7H' \
    --namespace 'Jk0F89xA' \
    --userId 'c3YVfaEN' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteChannel'
delete_file $TEMP_FILE

#- 36 adminUpdateContentS3
samples/cli/sample-apps Ugc adminUpdateContentS3 \
    --body '{"contentType": "trl0pTKZ", "fileExtension": "TXqzHuBM", "name": "YQSA2jz1", "preview": "ZOpdOjSy", "subType": "MddB41Ju", "tags": ["Mf7RUyBH"], "type": "Rj8IiRim"}' \
    --channelId 'RllHT6Dc' \
    --contentId '40vFFA6g' \
    --namespace 'pU7EW3x1' \
    --userId 'dCpm55gO' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateContentS3'
delete_file $TEMP_FILE

#- 37 adminUpdateContentDirect
update_status 0 'adminUpdateContentDirect (skipped deprecated)'

#- 38 adminDeleteContent
samples/cli/sample-apps Ugc adminDeleteContent \
    --channelId 'eqQIqcJV' \
    --contentId 'KmBM1J1I' \
    --namespace 'buTrrkbm' \
    --userId 'uT1whOqm' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteContent'
delete_file $TEMP_FILE

#- 39 adminGetContent
samples/cli/sample-apps Ugc adminGetContent \
    --namespace 'EnDXIWrB' \
    --userId 'PlSay46m' \
    --limit 'v71BAZAO' \
    --offset 'jtFJ2vmT' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetContent'
delete_file $TEMP_FILE

#- 40 adminDeleteAllUserContents
samples/cli/sample-apps Ugc adminDeleteAllUserContents \
    --namespace 'j7tT7TZH' \
    --userId 'WDdCkIsZ' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteAllUserContents'
delete_file $TEMP_FILE

#- 41 adminHideUserContent
samples/cli/sample-apps Ugc adminHideUserContent \
    --body '{"isHidden": false}' \
    --contentId 'ArWwPHcy' \
    --namespace 'FAdAtYci' \
    --userId 'LIgRwFRr' \
    >$TEMP_FILE 2>&1
update_status $? 'adminHideUserContent'
delete_file $TEMP_FILE

#- 42 adminGetAllGroups
samples/cli/sample-apps Ugc adminGetAllGroups \
    --namespace '0gwB9tz3' \
    --userId 'vp99XVlV' \
    --limit '8rK3tE6n' \
    --offset '0smip1tw' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetAllGroups'
delete_file $TEMP_FILE

#- 43 adminDeleteAllUserGroup
samples/cli/sample-apps Ugc adminDeleteAllUserGroup \
    --namespace '3L7cUd9p' \
    --userId 'qtv6JfPZ' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteAllUserGroup'
delete_file $TEMP_FILE

#- 44 adminGetGroup
samples/cli/sample-apps Ugc adminGetGroup \
    --groupId 'wcCVOXcV' \
    --namespace 'a80TmCwt' \
    --userId 'D2lAH01o' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetGroup'
delete_file $TEMP_FILE

#- 45 adminUpdateGroup
samples/cli/sample-apps Ugc adminUpdateGroup \
    --body '{"contents": ["6NdcBIgz"], "name": "rDyWpFBY"}' \
    --groupId 'GmmBawMy' \
    --namespace 'oKyNpdAa' \
    --userId 'sm8xwUfz' \
    >$TEMP_FILE 2>&1
update_status $? 'adminUpdateGroup'
delete_file $TEMP_FILE

#- 46 adminDeleteGroup
samples/cli/sample-apps Ugc adminDeleteGroup \
    --groupId 'OlQiZY4N' \
    --namespace 'bOQXJ7uO' \
    --userId 'TzNMvuq2' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteGroup'
delete_file $TEMP_FILE

#- 47 adminGetGroupContents
samples/cli/sample-apps Ugc adminGetGroupContents \
    --groupId 'tNl4CX4I' \
    --namespace 'jiK4DEUJ' \
    --userId 'RVK3l9Eb' \
    --limit '0R1XRb0R' \
    --offset 'H8vS1sme' \
    >$TEMP_FILE 2>&1
update_status $? 'adminGetGroupContents'
delete_file $TEMP_FILE

#- 48 adminDeleteAllUserStates
samples/cli/sample-apps Ugc adminDeleteAllUserStates \
    --namespace 'OlngrdTX' \
    --userId 'CzaPBtkZ' \
    >$TEMP_FILE 2>&1
update_status $? 'adminDeleteAllUserStates'
delete_file $TEMP_FILE

#- 49 searchChannelSpecificContent
samples/cli/sample-apps Ugc searchChannelSpecificContent \
    --channelId 'Mio4wcyh' \
    --namespace 'loVS3rYp' \
    --creator '8QtcEmCE' \
    --isofficial 'Vc75Ufey' \
    --limit 'pWjDNhzC' \
    --name 'L5sWS2qw' \
    --offset 'O763iEkl' \
    --orderby 'kzLm88Lp' \
    --sortby 'LuYRO3C5' \
    --subtype '5yHpwK2J' \
    --tags 'aqenDGn7' \
    --type 'a2NUplWi' \
    --userId 'Ljq06n6a' \
    >$TEMP_FILE 2>&1
update_status $? 'searchChannelSpecificContent'
delete_file $TEMP_FILE

#- 50 publicSearchContent
samples/cli/sample-apps Ugc publicSearchContent \
    --namespace '0rW8Efkp' \
    --creator 'aXtwYZJa' \
    --isofficial 'Q4WbwNms' \
    --limit 'FYetjEur' \
    --name 'H8eloJzN' \
    --offset 'KtRUaTz1' \
    --orderby 'ETdsmwzj' \
    --sortby 'kkn9oiQl' \
    --subtype '05g7cO3Z' \
    --tags 'Mb6Ojlo6' \
    --type 'DMNpP2qM' \
    --userId 'rTQ1Upjf' \
    >$TEMP_FILE 2>&1
update_status $? 'publicSearchContent'
delete_file $TEMP_FILE

#- 51 getFollowedContent
samples/cli/sample-apps Ugc getFollowedContent \
    --namespace 'U6wJhy1j' \
    --limit 'OVkkUlS7' \
    --offset '9527EZ25' \
    >$TEMP_FILE 2>&1
update_status $? 'getFollowedContent'
delete_file $TEMP_FILE

#- 52 getLikedContent
samples/cli/sample-apps Ugc getLikedContent \
    --namespace 'Ia8uCeZF' \
    --limit 'lLtEVpDA' \
    --offset 'EbA82jy7' \
    >$TEMP_FILE 2>&1
update_status $? 'getLikedContent'
delete_file $TEMP_FILE

#- 53 downloadContentByShareCode
samples/cli/sample-apps Ugc downloadContentByShareCode \
    --namespace '4lq0pDE5' \
    --shareCode 'xRwh5b45' \
    >$TEMP_FILE 2>&1
update_status $? 'downloadContentByShareCode'
delete_file $TEMP_FILE

#- 54 publicDownloadContentByContentID
samples/cli/sample-apps Ugc publicDownloadContentByContentID \
    --contentId 'ebpcM7Sc' \
    --namespace 'Ss3UOpAw' \
    >$TEMP_FILE 2>&1
update_status $? 'publicDownloadContentByContentID'
delete_file $TEMP_FILE

#- 55 addDownloadCount
samples/cli/sample-apps Ugc addDownloadCount \
    --contentId 'Ip9rRtn1' \
    --namespace 'PcCxdbum' \
    >$TEMP_FILE 2>&1
update_status $? 'addDownloadCount'
delete_file $TEMP_FILE

#- 56 updateContentLikeStatus
samples/cli/sample-apps Ugc updateContentLikeStatus \
    --body '{"likeStatus": false}' \
    --contentId 'YgOdEBWR' \
    --namespace 'QiW3KFfU' \
    >$TEMP_FILE 2>&1
update_status $? 'updateContentLikeStatus'
delete_file $TEMP_FILE

#- 57 publicDownloadContentPreview
samples/cli/sample-apps Ugc publicDownloadContentPreview \
    --contentId '8icH4081' \
    --namespace 'gRB1GyLf' \
    >$TEMP_FILE 2>&1
update_status $? 'publicDownloadContentPreview'
delete_file $TEMP_FILE

#- 58 getTag
samples/cli/sample-apps Ugc getTag \
    --namespace 'Lg4RYuEb' \
    --limit 'gUDEcJyI' \
    --offset 'vsPwOr0B' \
    >$TEMP_FILE 2>&1
update_status $? 'getTag'
delete_file $TEMP_FILE

#- 59 getType
samples/cli/sample-apps Ugc getType \
    --namespace 'mV5iFvfw' \
    --limit 'FjTSmIEq' \
    --offset 'oLyLeUGm' \
    >$TEMP_FILE 2>&1
update_status $? 'getType'
delete_file $TEMP_FILE

#- 60 getFollowedUsers
samples/cli/sample-apps Ugc getFollowedUsers \
    --namespace 'omGX9sXT' \
    --limit 'Z0v8pqLf' \
    --offset 'c5SwGnRe' \
    >$TEMP_FILE 2>&1
update_status $? 'getFollowedUsers'
delete_file $TEMP_FILE

#- 61 getChannels
samples/cli/sample-apps Ugc getChannels \
    --namespace 'UULDX4QU' \
    --userId 'Ibb5nh68' \
    --limit 'ZnyUtRvW' \
    --offset '9hNBSFTt' \
    >$TEMP_FILE 2>&1
update_status $? 'getChannels'
delete_file $TEMP_FILE

#- 62 createChannel
samples/cli/sample-apps Ugc createChannel \
    --body '{"name": "FrOmjkFr"}' \
    --namespace 'FVA8t0xF' \
    --userId '34Xpt6Zl' \
    >$TEMP_FILE 2>&1
update_status $? 'createChannel'
delete_file $TEMP_FILE

#- 63 deleteAllUserChannel
samples/cli/sample-apps Ugc deleteAllUserChannel \
    --namespace 'TTic0kr2' \
    --userId 'a0nI2oo7' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteAllUserChannel'
delete_file $TEMP_FILE

#- 64 updateChannel
samples/cli/sample-apps Ugc updateChannel \
    --body '{"name": "UHCJK5sp"}' \
    --channelId '0aCvIq3a' \
    --namespace 'HVYIlewL' \
    --userId 'RuHY83bG' \
    >$TEMP_FILE 2>&1
update_status $? 'updateChannel'
delete_file $TEMP_FILE

#- 65 deleteChannel
samples/cli/sample-apps Ugc deleteChannel \
    --channelId 'j0HTeeWX' \
    --namespace 'lIcRidqc' \
    --userId 'tDpygY0a' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteChannel'
delete_file $TEMP_FILE

#- 66 createContentDirect
update_status 0 'createContentDirect (skipped deprecated)'

#- 67 createContentS3
samples/cli/sample-apps Ugc createContentS3 \
    --body '{"contentType": "x476ED4M", "fileExtension": "MO9Tw2JH", "name": "0qhWIwHW", "preview": "TgzJFRYw", "subType": "6t1IKZLO", "tags": ["6V4Ode46"], "type": "QmCidgdp"}' \
    --channelId 'P7RTC587' \
    --namespace 'lmUmBziP' \
    --userId 'ZBnpOfkl' \
    >$TEMP_FILE 2>&1
update_status $? 'createContentS3'
delete_file $TEMP_FILE

#- 68 updateContentS3
samples/cli/sample-apps Ugc updateContentS3 \
    --body '{"contentType": "lxfq0Nsr", "fileExtension": "Sjw5Hog0", "name": "blM1d5MS", "preview": "tYGczLIN", "subType": "lEC0OEsE", "tags": ["3yzIsUP0"], "type": "NjluOrGZ"}' \
    --channelId 'TzsLW7Fj' \
    --contentId 'fs9nIkcZ' \
    --namespace '38fUEanj' \
    --userId 'KHbXfk1z' \
    >$TEMP_FILE 2>&1
update_status $? 'updateContentS3'
delete_file $TEMP_FILE

#- 69 updateContentDirect
update_status 0 'updateContentDirect (skipped deprecated)'

#- 70 deleteContent
samples/cli/sample-apps Ugc deleteContent \
    --channelId 'xdzxg0UX' \
    --contentId 'cRyHi3u8' \
    --namespace 'BzVWu1tO' \
    --userId 'mhUtCgcp' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteContent'
delete_file $TEMP_FILE

#- 71 publicGetUserContent
samples/cli/sample-apps Ugc publicGetUserContent \
    --namespace 'vGrEbcZU' \
    --userId 'DExH1tay' \
    --limit 'OGXIHzMR' \
    --offset 'jMCtOJsE' \
    >$TEMP_FILE 2>&1
update_status $? 'publicGetUserContent'
delete_file $TEMP_FILE

#- 72 deleteAllUserContents
samples/cli/sample-apps Ugc deleteAllUserContents \
    --namespace 'ijlrbpyy' \
    --userId 'EcQxVgJI' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteAllUserContents'
delete_file $TEMP_FILE

#- 73 updateScreenshots
samples/cli/sample-apps Ugc updateScreenshots \
    --body '{"screenshots": [{"description": "jMZqcWfM", "screenshotId": "l6dqrpD4"}]}' \
    --contentId 'tnc3ZRB3' \
    --namespace 'IkdtPfAJ' \
    --userId 'EomwenJv' \
    >$TEMP_FILE 2>&1
update_status $? 'updateScreenshots'
delete_file $TEMP_FILE

#- 74 uploadContentScreenshot
samples/cli/sample-apps Ugc uploadContentScreenshot \
    --body '{"screenshots": [{"contentType": "Q8grtQSv", "description": "6EcALcMI", "fileExtension": "png"}]}' \
    --contentId 'ms5bT51M' \
    --namespace '4yko8S0E' \
    --userId 'nGLvGvfu' \
    >$TEMP_FILE 2>&1
update_status $? 'uploadContentScreenshot'
delete_file $TEMP_FILE

#- 75 deleteContentScreenshot
samples/cli/sample-apps Ugc deleteContentScreenshot \
    --contentId 'SyCTyjj4' \
    --namespace 'mCaiuMGK' \
    --screenshotId 'OF5GJJoo' \
    --userId 'SXUl3YU3' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteContentScreenshot'
delete_file $TEMP_FILE

#- 76 updateUserFollowStatus
samples/cli/sample-apps Ugc updateUserFollowStatus \
    --body '{"followStatus": false}' \
    --namespace 'BABnOlxD' \
    --userId 'znICQVyq' \
    >$TEMP_FILE 2>&1
update_status $? 'updateUserFollowStatus'
delete_file $TEMP_FILE

#- 77 getGroups
samples/cli/sample-apps Ugc getGroups \
    --namespace 'Bg34WTtD' \
    --userId 'kn0rtn6t' \
    --limit '0Yx4z12E' \
    --offset 'aQ1rUQYC' \
    >$TEMP_FILE 2>&1
update_status $? 'getGroups'
delete_file $TEMP_FILE

#- 78 createGroup
samples/cli/sample-apps Ugc createGroup \
    --body '{"contents": ["NTiDX4jE"], "name": "3M2IsTHu"}' \
    --namespace '8QwNyOlX' \
    --userId 'fIWd0mcq' \
    >$TEMP_FILE 2>&1
update_status $? 'createGroup'
delete_file $TEMP_FILE

#- 79 deleteAllUserGroup
samples/cli/sample-apps Ugc deleteAllUserGroup \
    --namespace '5T4SUc7c' \
    --userId 'WfCKK6Di' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteAllUserGroup'
delete_file $TEMP_FILE

#- 80 getGroup
samples/cli/sample-apps Ugc getGroup \
    --groupId 'j1gFcenE' \
    --namespace 'MySPfhxB' \
    --userId 'enDiTiAq' \
    >$TEMP_FILE 2>&1
update_status $? 'getGroup'
delete_file $TEMP_FILE

#- 81 updateGroup
samples/cli/sample-apps Ugc updateGroup \
    --body '{"contents": ["FYmFKjaE"], "name": "Lmmll6oe"}' \
    --groupId 'xId1OKGU' \
    --namespace 'N2Uznd7u' \
    --userId 'Va7t14yv' \
    >$TEMP_FILE 2>&1
update_status $? 'updateGroup'
delete_file $TEMP_FILE

#- 82 deleteGroup
samples/cli/sample-apps Ugc deleteGroup \
    --groupId 'SYSV52bH' \
    --namespace 'ifCIf4ts' \
    --userId 'uu6Pkam6' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteGroup'
delete_file $TEMP_FILE

#- 83 getGroupContent
samples/cli/sample-apps Ugc getGroupContent \
    --groupId 'tFSYFt4Z' \
    --namespace 'xA2PzZFR' \
    --userId 'kBNlg6hn' \
    --limit '5qusKyZA' \
    --offset 'uV6uUvqM' \
    >$TEMP_FILE 2>&1
update_status $? 'getGroupContent'
delete_file $TEMP_FILE

#- 84 deleteAllUserStates
samples/cli/sample-apps Ugc deleteAllUserStates \
    --namespace '0lV6UZMl' \
    --userId 'EbxHNgJR' \
    >$TEMP_FILE 2>&1
update_status $? 'deleteAllUserStates'
delete_file $TEMP_FILE

delete_file 'tmp.dat'

exit $FAILED_COUNT