import { memo } from 'react';
import { useTranslation } from 'react-i18next';
import '../../../'
import DiceGrid from '../../../components/dice-grid';
import DicePlate from '../../../components/dice-plate';

import { formatThousand, classnames } from '../../utils';
import { useDiceGame } from './utils';

import classes from './styles.module.css';

export const DiceGameComponent: React.FC = () => {
  const { t } = useTranslation();
  const {
    amount,
    betState,
    rolledDices,
    needToShowResult,
    rolling,
    handleBet,
    handleResetBet,
    handleRoll,
    startNewSession,
  } = useDiceGame();
  return (
    <div data-testid="dice-game">
      <div className={classes['amount-section']}>
        <span>${formatThousand(amount)}</span>
      </div>
      <DicePlate disabled={rolling} items={rolledDices} onStart={handleRoll} />
      <DiceGrid
        betValues={betState}
        rolledDices={rolledDices}
        needToShowResult={needToShowResult}
        onBet={handleBet}
        onResetBet={handleResetBet}
      />
    </div>
  );
};

const DiceGame = memo(DiceGameComponent);
DiceGame.displayName = 'DiceGame';

export default DiceGame;
