using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace LP19_App
{
    public partial class frmMain : Form
    {
        private Classes.Connection oConnection;
        public frmMain()
        {
            InitializeComponent();
        }

        private void frmMain_Load(object sender, EventArgs e)
        {
            cbConnStatus.CheckState = CheckState.Unchecked;
            cbConnStatus.Text = Properties.Resources.sNoConnection;

            oConnection = new Classes.Connection();

            tConnInterval.Enabled = true;
        }

        private void tConnInterval_Tick(object sender, EventArgs e)
        {
            bool bResult = oConnection.establishConnection();

            Console.WriteLine(bResult);
        }
    }
}
