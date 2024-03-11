namespace LP19_App
{
    partial class frmMain
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.components = new System.ComponentModel.Container();
            this.tlpWidgets = new System.Windows.Forms.TableLayoutPanel();
            this.cbConnStatus = new System.Windows.Forms.CheckBox();
            this.tConnInterval = new System.Windows.Forms.Timer(this.components);
            this.tlpWidgets.SuspendLayout();
            this.SuspendLayout();
            // 
            // tlpWidgets
            // 
            this.tlpWidgets.ColumnCount = 1;
            this.tlpWidgets.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tlpWidgets.Controls.Add(this.cbConnStatus, 0, 0);
            this.tlpWidgets.Location = new System.Drawing.Point(12, 12);
            this.tlpWidgets.Name = "tlpWidgets";
            this.tlpWidgets.RowCount = 3;
            this.tlpWidgets.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50.54945F));
            this.tlpWidgets.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 49.45055F));
            this.tlpWidgets.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 347F));
            this.tlpWidgets.Size = new System.Drawing.Size(289, 439);
            this.tlpWidgets.TabIndex = 0;
            // 
            // cbConnStatus
            // 
            this.cbConnStatus.AutoCheck = false;
            this.cbConnStatus.AutoSize = true;
            this.cbConnStatus.Dock = System.Windows.Forms.DockStyle.Fill;
            this.cbConnStatus.Font = new System.Drawing.Font("Microsoft Sans Serif", 11.25F, System.Drawing.FontStyle.Regular, System.Drawing.GraphicsUnit.Point, ((byte)(0)));
            this.cbConnStatus.Location = new System.Drawing.Point(3, 3);
            this.cbConnStatus.Name = "cbConnStatus";
            this.cbConnStatus.Size = new System.Drawing.Size(283, 40);
            this.cbConnStatus.TabIndex = 0;
            this.cbConnStatus.UseVisualStyleBackColor = true;
            // 
            // tConnInterval
            // 
            this.tConnInterval.Tick += new System.EventHandler(this.tConnInterval_Tick);
            // 
            // frmMain
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(313, 463);
            this.Controls.Add(this.tlpWidgets);
            this.Name = "frmMain";
            this.Text = "Controller";
            this.Load += new System.EventHandler(this.frmMain_Load);
            this.tlpWidgets.ResumeLayout(false);
            this.tlpWidgets.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private System.Windows.Forms.TableLayoutPanel tlpWidgets;
        private System.Windows.Forms.CheckBox cbConnStatus;
        private System.Windows.Forms.Timer tConnInterval;
    }
}

