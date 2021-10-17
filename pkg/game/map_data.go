package game

var MapIndex map[int]string

func init() {
	MapIndex = make(map[int]string)
	MapIndex[0] = Map0
	MapIndex[1] = Map1
	MapIndex[2] = Map2
	MapIndex[3] = Map3
	MapIndex[4] = Map4
	MapIndex[5] = Map5
	MapIndex[6] = Map6
	MapIndex[7] = Map7
	MapIndex[8] = Map8
	MapIndex[9] = Map9
}

const (
	Map0 = `########################
#Every #body is!#~a    #
####   ###### ######## #
#!genius. But if# you  #
#not!###  #######  #####
# judge# a fish by its #
# ####  ~ #### ### #####
#  ability to   #climb #
# a  #tree,# it # will #
#live# ### # its### ####
#   ~#whole#life#     ~#
###when##  ### you###  #
  # believing that# it##
  ######### ##### is### 
          #  stupid   # 
          ############# `

	Map1 = `#######################
# What... is the air- #
#  ## ####~  #### ##~ #
#  #speed#   # of  #  #
# ~#### ##  ~## ####  #
# an unladen swallow? #
#What..### do ## you  #
#   ~## mean?   ## An #
# African or European #
#swallow?#####~Huh?   #
# I... I don't know   #
##  that.. #~*thrown*##
 ### *into*# *the* ### 
   ### *volcano* ###   
     #############     `
	// 注意后面加空格哦

	Map2 = `#############
#Do~  #not~ #
# try #~and #
#bend #the~ #
####spoon####
# ~Instead, #
#####try#####
#to realize #
#the#truth. #
#   ####    #
#   There~is#
#~no#spoon.~#
#   #       #
#############`

	Map3 = `###################################
# Will all great # Neptune's Ocean#
# wash this~blood# clean~from my~~#
# hand?~No, this # my hand~will~~~#
# rather the ##### multitudinous###
#seas incarnadine# ~  making the  #
# green~one red. # None of woman  #
# born shall harm# Macbeth.~      #
#Macbeth ~~ shall#never vanquish'd#
# be until Great # Birnam wood to #
# high~Dunsinane # hill shall~come#
# against him.#### By the pricking#
# of my~thumbs,  #   something#####
#    wicked ~this way~  comes.#   #
###################################`

	Map4 = `               ###                         ###           
              ## #    ###           ###    # ##          
            ### ##    # ##         ## #    ## ###        
          ###He##     #is###########  #     ##is###      
        ###the##      #   hero that   #      ##   ###    
       ##    ##       #Gotham deserves#       ## but##   
      ## not #        # the one it    #        #needs##  
     ## right###    ###now, so we'll  ###    ###hunt  ## 
    ## him.    ######   Because he can  ###### take it.##
    #  Because he's not a hero. He's a silent guardian, #
    ## a watchful protector, a Dark Knight. It is not  ##
     ##  who we are undearneath, but what we do that  ## 
      ##  that  ###  ### defines us. ###  ### Bruce, ##  
       ##why do## ## # ## we fall?  ## #2## ## pick ##   
        ##  our#   ###  ##selves up##  ###   #again##    
         ###Kil##        ##ling is##        ##not###     
           ###ju##        ##stice##        ##. ###       
             ##  #         ##TDK##         #  ##         
              ####          ##!##          ####          
                             ###                         `

	Map5 = `######################
#That it should come #
#to##### #######   ###
# this!#    #Frailty,#
###   ##thy #   ######
#name~is!####   #alas#
# poor yorick! This  #
####above#####   #####
   #all  ##### to    #
####thine be~true## ~#
#  #     #####   ## ~#
#  #Some ###thing##is#
#  ####   #    #######
#rotten  ~#in  #     #
#######   ###  #######
      # Denmark#      
      ##########      `

	Map6 = `##################      ##############     
#Hey, Doc, we#   #      #   ~Roads?  #     
# better back#up.#      #where we're #     
# We don't have  #      # going, we ~#     
#enough~road to  #      #don't###    #     
# get up#to 88.  #      #~need#roads!#     
##################      ##############     
                                           
                                           
###################     ###################
#Doc! Doc! No! No!#     # What about all  #
# You're alive!!! #     #that talk about  #
#Bulletproof~vest?#     #screwing up#the  #
#How did you know?#     # space-time#######
#  I never got a  #     #####continuum?   #
#~chance to tell  #     # Well, I figured,#
####### you.~######     #~what the hell?  #
      ########          ###################`

	Map7 = `####################
#That#Terminator is#
#out # there. It~~~#
# can't be reasoned#
# with. It doesn't #
#feel ### pity or###
# remorse ### or ###
# fear ~ and it ~  #
# absolutely ~ will#
#not ## stop. Ever.#
# Until ~ you ~ are#
#     ~ dead ~     #
####################`

	Map8 = `                       ###########             
                    ####Waka waka####          
                 #### waka waka waka####       
               ###waka~waka ~waka~ waka###     
             ###"Video games don't affect###   
           ###kids. If Pacman had affected ### 
          ## us as kids,~we'd all be running ##
         ## in~darkened rooms, ~ munching  ### 
        ## magic pills and listening to  ###   
       ## repetitive electronic music" ###     
       ##~~~~~- Marcus Brigstocke    ###       
       ##     ~~(CEO of Nintendo)~ ###         
       ##Waka waka waka waka waka wak###       
       ##~ waka~waka waka waka waka wak###     
        ## waka waka~waka waka~waka waka~###   
         ## waka~waka waka~waka waka~waka ~### 
          ##waka waka~waka waka waka waka wak##
           ###waka waka waka~waka waka~waka### 
             ### waka~waka waka waka~waka###   
               ### waka waka~waka waka ###     
                 #### waka waka waka####       
                    ####  WAKA'  ####          
                       ###########             `

	Map9 = `##################
#S~ay~~some~thing#
#c~oo~l w~hen~you#
#thr~ow~it!~ONE,~#
#T~W~O!~~THR~EE !#
#S~OM~E THI~~~NG~#
#~~C~~O~~~O~~L! ~#
##################`

	MapForTest = `#  b~a  #`
)
