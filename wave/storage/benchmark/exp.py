import os
import time

UPDATES=[0,100]#5,50,95]
THREADS=[1,4,8,12,16,20,24,28]
RPT=10
DURATION=30
LOAD=100
REQUEST_SIZE=100

SERVER_DIR="/opt/src/wave/storage/vldmstorage3/server/init"

SETUP_CMD = "./setup.sh"
STOP_CMD = "./stop.sh"
# out_{rps}_{update}_r{run}
EXP_CMD = "go run . {} {} {} {} {} {} > out_{}_{}_r{} 2>&1"
CP_CMD = "cp {} {}" 

def run_exp(t, rps, d, u, ls, rs, number):
  os.system(SETUP_CMD)
  time.sleep(1) 
  os.system(EXP_CMD.format(t, rps, d, u, ls, rs, rps, u, number))
  os.system(STOP_CMD) 
  os.system(CP_CMD.format("{}/server_logs".format(SERVER_DIR), "server_log_{}_{}_r{}".format(rps, u, number)))

if __name__=="__main__":
  for exp in range(3):
    for t in THREADS:
      for u in UPDATES:
        print("Running with {} threads, {}% update".format(t, u))
        run_exp(t, t*RPT, DURATION, u, LOAD, REQUEST_SIZE, exp)
        time.sleep(2)
