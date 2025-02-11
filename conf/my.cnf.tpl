[mysqld]
datadir=/var/lib/mysql
#socket=/var/lib/mysql/mysql.sock
symbolic-links=0
log-error=/var/lib/mysql/mysqld.log
pid-file=/var/run/mysqld/mysqld.pid

#开启GTID,必须开启
gtid_mode = ON
#强制GTID的一致性
enforce_gtid_consistency = ON
#binlog格式,MGR要求必须是ROW,不过就算不是MGR,也最好用
binlog_format=ROW
#server-id必须是唯一的
server-id = {{ SERVER_ID }}
#MGR使用乐观锁,所以官网建议隔离级别是RC,减少锁粒度
transaction_isolation = READ-COMMITTED
#因为集群会在故障恢复时互相检查binlog的数据,
#所以需要记录下集群内其他服务器发过来已经执行过的binlog,按GTID来区分是否执行过.
log-slave-updates = 1
#binlog校验规则,5.6之后的高版本是CRC32,低版本都是NONE,但是MGR要求使用NONE
binlog_checksum = NONE
#基于安全的考虑,MGR集群要求复制模式要改成slave记录到表中,不然就报错
master_info_repository = TABLE
#同上配套
relay_log_info_repository = TABLE
#组复制设置
#记录事务的算法,官网建议设置该参数使用 XXHASH64 算法
transaction_write_set_extraction = XXHASH64
#相当于此GROUP的名字,是UUID值,不能和集群内其他GTID值的UUID混用,可用uuidgen来生成一个新的,
#主要是用来区分整个内网里边的各个不同的GROUP,而且也是这个group内的GTID值的UUID
loose-group_replication_group_name = '{{ GROUP_NAME }}'
#IP地址白名单,默认只添加127.0.0.1,不会允许来自外部主机的连接,按需安全设置
#loose-group_replication_ip_whitelist = '127.0.0.1/8,192.168.32.0/24'
#是否随服务器启动而自动启动组复制,不建议直接启动,怕故障恢复时有扰乱数据准确性的特殊情况
loose-group_replication_start_on_boot = OFF
#本地MGR的IP地址和端口，host:port,是MGR的端口,不是数据库的端口
loose-group_replication_local_address = '{{ LOCAL_ADDRESS }}'
#需要接受本MGR实例控制的服务器IP地址和端口,是MGR的端口,不是数据库的端口
loose-group_replication_group_seeds = '{{ GROUP_SEEDS }}'
#开启引导模式,添加组成员，用于第一次搭建MGR或重建MGR的时候使用,只需要在集群内的其中一台开启,
loose-group_replication_bootstrap_group = OFF
#是否启动单主模式，如果启动，则本实例是主库，提供读写，其他实例仅提供读,如果为off就是多主模式了
loose-group_replication_single_primary_mode = ON
#多主模式下,强制检查每一个实例是否允许该操作,如果不是多主,可以关闭
#loose-group_replication_enforce_update_everywhere_checks = off
# gtid 不一致情况下是否允许加入，默认关闭
#loose-group_replication_allow_local_disjoint_gtids_join=OFF

max_connections=10000
open_files_limit = 65535
default-time-zone='+8:00'
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
default_authentication_plugin = mysql_native_password
log_error_suppression_list = 'MY-013360'