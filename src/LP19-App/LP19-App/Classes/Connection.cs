using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.IO.Ports;

namespace LP19_App.Classes
{
    class Connection
    {
        private SerialPort serialPort;        
        
        public Connection()
        {
            serialPort = new SerialPort("COM6", 230400, Parity.None, 8, StopBits.One);
            serialPort.ReadTimeout = 5000;
        }

        public bool establishConnection()
        {
            bool bReturn = false;
            byte[] buf = new byte[1];

            if (serialPort.IsOpen)
            {
                serialPort.Read(buf, 0, buf.Length);
                byte message = buf[0];
                Console.WriteLine(message);

                bReturn = true;
            } else
            {
                bReturn = false;
            }

            return bReturn;
        }
    }
}
