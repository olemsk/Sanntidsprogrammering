with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure exercise7 is

    Count_Failed    : exception;    -- Exception to be raised when counting fails
    Gen             : Generator;    -- Random number generator
	
    protected type Transaction_Manager (N : Positive) is
        entry Finished;
        entry Wait_Until_Aborted;
       -- function Commit return Boolean;
        procedure Signal_Abort;
    private
        Finished_Gate_Open  : Boolean := False;
        Aborted             : Boolean := False;
        Num_Informed		: Natural := 0; 
    end Transaction_Manager;
    protected body Transaction_Manager is
        entry Finished when Finished_Gate_Open or Finished'Count = N is
        begin
			Finished_Gate_Open := True;
			--Put_Line("Heluuuu");
			if Finished'Count = 0 then
				Finished_Gate_Open := False;
			end if;
        end Finished;
        
    	entry Wait_Until_Aborted when Aborted is
    	begin
    	    Num_Informed := Num_Informed +1; 
    		if Num_Informed = 3 then
			Aborted := False;
			Num_Informed := 0;
			end if;
		
  
    		
    		--Manager.Signal_Abort;
    	end Wait_Until_Aborted;	

        procedure Signal_Abort is
        begin
            Aborted := True;
        end Signal_Abort;

        --function Commit return Boolean is
       -- begin
          --  return Should_Commit;
        --end Commit;
        
    end Transaction_Manager;



    
    function Unreliable_Slow_Add (x : Integer) return Integer is
    Error_Rate : Constant := 0.15;  -- (between 0 and 1)
    begin
        if Random(Gen)>0.15 then
        	delay Duration(4);
        	return x+10;
        else
        	raise Count_Failed;
        	delay Duration(0.5);
        end if;
        -------------------------------------------
        -- PART 1: Create the transaction work here
        -------------------------------------------
    end Unreliable_Slow_Add;




    task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
    task body Transaction_Worker is
        Num         : Integer   := Initial;
        Prev        : Integer   := Num;
        Round_Num   : Integer   := 0;
    begin
        Put_Line ("Worker" & Integer'Image(Initial) & " started");

        loop
            Put_Line ("Worker" & Integer'Image(Initial) & " started round" & Integer'Image(Round_Num));
            Round_Num := Round_Num + 1;

			select
    			Manager.Wait_Until_Aborted;   -- eg. X.Entry_Call;
    			-- code that is run when the triggering_alternative has triggered
    			--   (forward ER code goes here)
    			Num := Num +5;
    			Put_Line ("  Worker" & Integer'Image(Initial) & " correcting" & Integer'Image(Num));

			then abort
    			--abortable_part
    			-- code that is run when nothing has triggered
    			--   (main functionality)
    			begin
    			Num := Unreliable_Slow_Add(Num);
    			exception
				when Count_Failed =>
						Manager.Signal_Abort;
			
				end;
				Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));

				Manager.Finished;		
			end select;
			
			
			
			
			

            --------------------------------------
            -- PART 2: Do the transaction work here             
            ---------------------------------------
            


            
              --  Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));

                -------------------------------------------
                -- PART 2: Roll back to previous value here
                -------------------------------------------
            

            Prev := Num;
            delay 0.5;

        end loop;
    end Transaction_Worker;

    Manager : aliased Transaction_Manager (3);

    Worker_1 : Transaction_Worker (0, Manager'Access);
    Worker_2 : Transaction_Worker (1, Manager'Access);
    Worker_3 : Transaction_Worker (2, Manager'Access);

begin
    Reset(Gen); -- Seed the random number generator
end exercise7;



